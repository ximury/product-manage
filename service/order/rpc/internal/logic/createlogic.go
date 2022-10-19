package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"product/service/order/model"
	"product/service/order/rpc/internal/svc"
	"product/service/order/rpc/pb/order"
	"product/service/product/rpc/pb/product"
	"product/service/user/rpc/pb/userclient"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 查询用户是否存在
	/*** 执行顺序
	1. /user/rpc/user/user.go: type User interface{}
	2. /user/rpc/user/user.go: func (m *defaultUser) GetUser()
	3. user/rpc/pb/userclient/user_grpc.pb.go: type UserClient interface[}
	4. user/rpc/pb/userclient/user_grpc.pb.go: func (c *userClient) GetUser()
	*/
	_, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.GetUserRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询产品是否存在
	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})
	if err != nil {
		return nil, err
	}
	// 判断产品库存是否充足
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "产品库存不足")
	}

	newOrder := model.Order{
		Uid:    uint64(in.Uid),
		Pid:    uint64(in.Pid),
		Amount: uint64(in.Amount),
		Status: 0,
	}
	// 创建订单
	res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	lastInsertId, err := res.LastInsertId()
	newOrder.Id = uint64(lastInsertId)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 更新产品库存
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     productRes.Id,
		Name:   productRes.Name,
		Desc:   productRes.Desc,
		Stock:  productRes.Stock - 1,
		Amount: productRes.Amount,
		Status: productRes.Status,
	})
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{
		Id: int64(newOrder.Id),
	}, nil
}
