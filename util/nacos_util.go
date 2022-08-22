package util

import (
	"errors"
	"fmt"
	"github.com/juqiukai/glog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"strconv"
	"strings"
)

const (
	CONF_KEY_MQ_ADDR    = "mq.addr"
	CONF_KEY_REDIS_ADDR = "redis.addr"
)

var nacosUtil = new(NacosUtil)

type NacosUtil struct {
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
	configMap    map[string]string
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
		return err
	}

	// Create config client for dynamic configuration
	cc, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if nil != err {
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
		return nil, err
	}
	if "" == content {
		return nil, errors.New(fmt.Sprintf("not found content, dataId=%s", dataId))
	}

	var confMap map[string]string
	err = yaml.Unmarshal([]byte(content), &confMap)
	if nil != err {
		return nil, err
	}
	u.configMap = confMap
	return confMap, nil
}

//
//  GetStringValue
//  @Description: 获取string值
//  @receiver u
//  @param key
//  @return string
//
func (u *NacosUtil) GetStringValue(key string) string {
	if u.configMap == nil {
		return ""
	}
	return u.configMap[key]
}

func (u *NacosUtil) DeRegister(port uint64) (bool, error) {
	ip, _ := GetLocalIp()
	return u.namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: "a.b.c",
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
//  getMQAddr
//  @Description: 获取mq地址
//  @receiver u
//  @return string
//
func (u *NacosUtil) GetMQAddr() string {
	return u.configMap[CONF_KEY_MQ_ADDR]
}

//
//  getRedisAddr
//  @Description: 获取redis地址
//  @receiver u
//  @return string
//
func (u *NacosUtil) GetRedisAddr() string {
	return u.configMap[CONF_KEY_REDIS_ADDR]
}
