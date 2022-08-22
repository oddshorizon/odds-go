package util

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

//
//  NewRedisClient
//  @Description:
//  @param confMap
//  @return *redis.Client
//  @return error
//
func NewRedisClient(confMap map[string]string) (*redis.Client, error) {
	addr := confMap["redis.addr"]
	if "" == addr {
		return nil, errors.New("conf - redis.addr='' ")
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}), nil
}
