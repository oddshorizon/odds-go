package test

import (
	"context"
	"github.com/oddshorizon/odds-go/util"
	"testing"
)

func TestRedisClusterGet(t *testing.T) {
	rdb, err := util.NewRedisClusterClient("192.168.0.55:7006", "123456")
	if nil != err {
		t.Fatalf("connect redis fail - %v", err)
	}
	val, err := rdb.Get(context.Background(), "name").Result()
	if nil != err {
		t.Fatalf("get redis val fail - %v", err)
	}
	t.Logf("name=%s", val)
}
