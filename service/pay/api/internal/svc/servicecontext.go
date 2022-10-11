package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"product/service/pay/api/internal/config"
	"product/service/pay/rpc/payclient"
)

type ServiceContext struct {
	Config config.Config

	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
