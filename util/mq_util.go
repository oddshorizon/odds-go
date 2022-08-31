package util

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"strings"
)

func NewMQProducer(mqAddr string) (rocketmq.Producer, error) {
	if "" == mqAddr {
		return nil, errors.New("conf - redis.addr='' ")
	}
	addrArr := strings.Split(mqAddr, ",")
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(addrArr)),
		producer.WithRetry(2),
	)
	return p, err
}

func NewMQPushConsumer(mqAddr, groupName string) (rocketmq.PushConsumer, error) {
	if "" == mqAddr {
		return nil, errors.New("conf - mq.addr='' ")
	}
	addrArr := strings.Split(mqAddr, ",")
	p, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(groupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(addrArr)),
	)
	return p, err
}

func NewMQAdmin(mqAddr string) (admin.Admin, error) {
	if "" == mqAddr {
		return nil, errors.New("conf - redis.addr='' ")
	}
	addrArr := strings.Split(mqAddr, ",")
	p, err := admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(addrArr)),
	)
	return p, err
}
