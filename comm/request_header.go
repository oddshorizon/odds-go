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
}

//
//  ParseRequestHeader
//  @Description: 解析出rpc请求头信息
//  @param ctx
//  @return *RequestHeader
//
func ParseRequestHeader(ctx context.Context) *RequestHeader {
	md, _ := metadata.FromIncomingContext(ctx)
	domainId := util.GetValFromMd(md, DOMAIN_ID)
	fromUid := util.GetValFromMd(md, USER_ID)
	tenantId := util.GetValFromMd(md, TENANT_ID)
	return &RequestHeader{DomainId: domainId, FromUid: fromUid, TenantId: tenantId}
}

func ParseRequestHeaderExt(ctx context.Context) (*RequestHeader, metadata.MD) {
	md, _ := metadata.FromIncomingContext(ctx)
	domainId := util.GetValFromMd(md, DOMAIN_ID)
	fromUid := util.GetValFromMd(md, USER_ID)
	tenantId := util.GetValFromMd(md, TENANT_ID)
	return &RequestHeader{DomainId: domainId, FromUid: fromUid, TenantId: tenantId}, md
}
