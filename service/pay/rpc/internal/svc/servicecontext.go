package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"product/service/order/rpc/orderclient"
	"product/service/pay/model"
	"product/service/pay/rpc/internal/config"
	"product/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel

	UserRpc  user.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,

		PayModel: model.NewPayModel(conn, c.CacheRedis),

		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
