package util

import (
	"errors"
	"github.com/go-redis/redis/v8"
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
func NewRedisClient(redisAddr string) (*redis.Client, error) {
	if redisAddr == "" {
		glog.Errorf("conf - redis.addr='' ")
		return nil, errors.New("redis.addr not allow blank")
	}
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	}), nil
}

//
//  NewRedisClusterClient
//  @Description:
//  @param redisAddrs  多个用 逗号分隔
//  @param password
//  @return *redis.ClusterClient
//  @return error
//
func NewRedisClusterClient(redisAddrs string, password string) (*redis.ClusterClient, error) {
	if "" == redisAddrs {
		glog.Errorf("conf - redis.addrs is blank ")
		return nil, errors.New("redis.addr not allow blank")
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(redisAddrs, ","),
		Password: password,
	})
	return rdb, nil
}
