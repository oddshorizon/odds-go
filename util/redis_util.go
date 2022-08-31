package util

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/juqiukai/glog"
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
