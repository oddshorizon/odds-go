package util

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

//
//  GetValFromMd
//  @Description: 从metadata中根据key查询对应的val，如果不存在，则返回空串
//  @param md
//  @param key
//  @return string
//
func GetValFromMd(md metadata.MD, key string) string {
	arr := md.Get(key)
	if nil == arr || len(arr) == 0 {
		return ""
	}
	return arr[0]
}

//
//  getClientAddr
//  @Description: 获取客户端地址
//  @receiver s
//  @param ctx
//  @return string
//
func GetClientAddr(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	var clientAddr string
	if ok {
		clientAddr = p.Addr.String()
	}
	return clientAddr
}
