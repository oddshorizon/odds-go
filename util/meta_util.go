package util

import "google.golang.org/grpc/metadata"

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
