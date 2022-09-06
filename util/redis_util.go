package util

import (
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/juqiukai/glog"
	"strings"
)

//
//  NewRedisClient
//  @Description:
//  @param confMap
//  @return *redis.Client
//  @return error
//
func NewRedisClient(redisAddr string, poolSize int) (*redis.Client, error) {
	if redisAddr == "" {
		glog.Errorf("conf - redis.addr='' ")
		return nil, errors.New("redis.addr not allow blank")
	}
	if poolSize <= 0 {
		poolSize = 1000
	}
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
		PoolSize: poolSize,
	}), nil
}

//
//  NewRedisClusterClient
//  @Description:
//  @param redisAddrs  多个用 逗号分隔
//  @param password
//  @param poolSize  连接池大小
//  @return *redis.ClusterClient
//  @return error
//
func NewRedisClusterClient(redisAddrs, password string, poolSize int) (*redis.ClusterClient, error) {
	if "" == redisAddrs {
		glog.Errorf("conf - redis.addrs is blank ")
		return nil, errors.New("redis.addr not allow blank")
	}
	if poolSize <= 0 {
		poolSize = 1000
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(redisAddrs, ","),
		Password: password,
		PoolSize: poolSize,
	})
	return rdb, nil
}
