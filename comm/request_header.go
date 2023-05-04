package comm

import (
	"context"
	"github.com/oddshorizon/odds-go/util"
	"google.golang.org/grpc/metadata"
)

const (
	USER_ID   = "user_id"
	TENANT_ID = "tenant_id"
	DOMAIN_ID = "domain_id"
)

type RequestHeader struct {
	DomainId string
	FromUid  string
	TenantId string
	SelfMap  map[string]string // 自定义的header信息
}

func NewRequestHeader() *RequestHeader {
	header := new(RequestHeader)
	header.SelfMap = make(map[string]string)
	return header
}

// SetValue
//
//	@Description: 设置header信息
//	@receiver header
//	@param key
//	@param val
func (header *RequestHeader) SetValue(key string, val string) {
	header.SelfMap[key] = val
}

// GetValue
//
//	@Description: 获取header信息
//	@receiver header
//	@param key
//	@return string
func (header *RequestHeader) GetValue(key string) string {
	return header.SelfMap[key]
}

// ParseRequestHeader
//
//	@Description: 解析出rpc请求头信息
//	@param ctx
//	@return *RequestHeader
func ParseRequestHeader(ctx context.Context) *RequestHeader {
	header, _ := ParseRequestHeaderExt(ctx)
	return header
}

// ParseRequestHeaderExt
//
//	@Description: 解析出rpc请求头信息
//	@param ctx
//	@return *RequestHeader
//	@return metadata.MD
func ParseRequestHeaderExt(ctx context.Context) (*RequestHeader, metadata.MD) {
	md, _ := metadata.FromIncomingContext(ctx)
	domainId := util.GetValFromMd(md, DOMAIN_ID)
	fromUid := util.GetValFromMd(md, USER_ID)
	tenantId := util.GetValFromMd(md, TENANT_ID)
	header := NewRequestHeader()
	header.DomainId = domainId
	header.FromUid = fromUid
	header.TenantId = tenantId
	return header, md
}
