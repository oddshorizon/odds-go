package util

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/oddshorizon/glog"
	"gopkg.in/yaml.v2"
	"strconv"
	"strings"
)

const (
	CONF_KEY_ROCKETMQ_ADDR = "rocketmq.addr"
	CONF_KEY_REDIS_ADDR    = "redis.addr"
)

var nacosUtil = new(NacosUtil)

type NacosUtil struct {
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
	configMap    map[string]string
	configId     string
}

func GetNacosUtil() *NacosUtil {
	return nacosUtil
}

//
//  LaunchNacosClients
//  @Description: 启动naco连接客户端
//  @param namespaceId
//  @param app
//  @param serverAddrs
//  @return error
//
func (u *NacosUtil) LaunchNacosClients(namespaceId, app string, serverAddrs string) error {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              fmt.Sprintf("/tmp/logs/%s", app),
		CacheDir:            fmt.Sprintf("/tmp/data/%s", app),
		LogLevel:            "info",
	}

	serverConfigs := []constant.ServerConfig{}
	arr := strings.Split(serverAddrs, ",")
	for i := range arr {
		hostPort := strings.Trim(arr[i], " ")
		if hostPort == "" {
			continue
		}
		hpArr := strings.Split(hostPort, ":")
		ip := strings.Trim(hpArr[0], " ")
		port := strings.Trim(hpArr[1], " ")
		intPort, _ := strconv.ParseInt(port, 10, 64)
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr:      ip,
			ContextPath: "/nacos",
			Port:        uint64(intPort),
			Scheme:      "http",
		})
	}

	// Create naming client for service discovery
	nc, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if nil != err {
		glog.Errorf("create nacos naming client fail - namespaceId=%s, app=%s, nacosServerAddrs=%s, err=%v", namespaceId, app, serverAddrs, err)
		return err
	}

	// Create config client for dynamic configuration
	cc, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if nil != err {
		glog.Errorf("create nacos config client fail - namespaceId=%s, app=%s, nacosServerAddrs=%s, err=%v", namespaceId, app, serverAddrs, err)
		return err
	}

	u.namingClient = nc
	u.configClient = cc
	return nil
}

//
//  LoadConfig
//  @Description: 加载配置
//  @param dataId
//  @return error
//
func (u *NacosUtil) LoadConfig(sysName string) (map[string]string, error) {
	dataId := fmt.Sprintf("%s.yaml", sysName)
	content, err := u.configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  "DEFAULT_GROUP",
	})
	if nil != err {
		glog.Errorf("load nacos config fail - dataId=%s", dataId)
		return nil, err
	}
	if "" == content {
		glog.Errorf("load nacos config content is blank - dataId=%s", dataId)
		return nil, errors.New(fmt.Sprintf("not found content, dataId=%s", dataId))
	}

	var confMap map[string]string
	err = yaml.Unmarshal([]byte(content), &confMap)
	if nil != err {
		glog.Errorf("unmarshal config content fail - content=%s, err=%v", content, err)
		return nil, err
	}
	u.configMap = confMap
	u.configId = dataId
	u.PrintConfigInfo()
	return confMap, nil
}

func (u *NacosUtil) PrintConfigInfo() {
	glog.Infof("load nacos conf success - %s ", u.configId)
	if nil != u.configMap {
		for k, v := range u.configMap {
			glog.Infof("%s: %s", k, v)
		}
	}
}

//
//  GetStringValue
//  @Description: 获取string类型的值
//  @receiver u
//  @param key
//  @param defaultValue
//  @return string
//
func (u *NacosUtil) GetStringValue(key, defaultValue string) string {
	retVal := defaultValue
	if u.configMap != nil {
		val, ok := u.configMap[key]
		if ok {
			retVal = val
		}
	}
	return retVal
}

func (u *NacosUtil) GetInt64Value(key string, defaultValue int64) int64 {
	retVal := defaultValue
	if u.configMap != nil {
		val, ok := u.configMap[key]
		if ok {
			int64Val, err := strconv.ParseInt(val, 10, 64)
			if nil == err {
				retVal = int64Val
			}
		}
	}
	return retVal
}

//
//  GetBoolValue
//  @Description: 获取bool值
//  @receiver u
//  @param key
//  @param defaultValue
//  @return bool
//
func (u *NacosUtil) GetBoolValue(key string, defaultValue bool) bool {
	retVal := defaultValue
	if u.configMap != nil {
		val, ok := u.configMap[key]
		if ok {
			if "true" == val {
				retVal = true
			} else if "false" == val {
				retVal = false
			}
		}
	}
	return retVal
}

//
//  DeRegister
//  @Description: 取消注册
//  @receiver u
//  @param port
//  @param serviceName
//  @return bool
//  @return error
//
func (u *NacosUtil) DeRegister(port uint64, serviceName string) (bool, error) {
	ip, _ := GetLocalIp()
	return u.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
	})
}

//
//  Register
//  @Description: 注册微服务
//  @receiver u
//  @param port
//  @param serviceName
//  @return bool
//  @return error
//
func (u *NacosUtil) Register(port uint64, serviceName string) (bool, error) {
	ip, _ := GetLocalIp()

	ok, err := u.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		ServiceName: serviceName,
		Ephemeral:   true,
	})
	if ok {
		glog.Infof("register service success - serviceName=%s, addr=%s:%d", serviceName, ip, port)
	}
	return ok, err
}

//
//  GetMQAddr
//  @Description: 获取rocketmq地址
//  @receiver u
//  @param defaultMQAddr
//  @return string
//
func (u *NacosUtil) GetRocketMQAddr(defaultMQAddr string) string {
	return u.GetStringValue(CONF_KEY_ROCKETMQ_ADDR, defaultMQAddr)
}

//
//  GetRedisAddr
//  @Description: 获取
//  @receiver u
//  @param defaultRedisAddr
//  @return string
//
func (u *NacosUtil) GetRedisAddr(defaultRedisAddr string) string {
	return u.GetStringValue(CONF_KEY_REDIS_ADDR, defaultRedisAddr)
}

//
//  GetServiceAddr
//  @Description: 获取微服务内网域名
//  @receiver u
//  @param serviceName
//  @param originalDomain
//  @return string
//
func (u *NacosUtil) GetServiceAddr(serviceName, originalAddr string) string {
	retAddr := originalAddr
	if "" != serviceName {
		if !strings.HasSuffix(serviceName, ".addr") {
			serviceName = fmt.Sprintf("%s.addr", serviceName)
		}
		addr, ok := u.configMap[serviceName]
		if ok {
			retAddr = addr
		}
	}
	return retAddr
}
