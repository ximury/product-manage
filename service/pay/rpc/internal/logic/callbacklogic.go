package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"product/service/order/rpc/pb/order"
	"product/service/pay/model"
	"product/service/user/rpc/user"

	"product/service/pay/rpc/internal/svc"
	"product/service/pay/rpc/pb/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Callback 添加支付回调逻辑
//支付流水回调流程：通过调用 user rpc 服务查询验证用户是否存在，
//再通过调用 order rpc 服务查询验证订单是否存在，
//然后通过查询库判断此订单支付流水是否存在，以及回调支付金额和库中流水支付金额是否一致，
//最后更新支付流水状态和通过调用 order rpc 服务更新订单状态
func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 查询支付是否存在
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 支付金额与订单金额不符
	resAmount := res.Amount
	if in.Amount != int64(resAmount) {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	res.Source = uint64(in.Source)
	res.Status = uint64(in.Status)

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CallbackResponse{}, nil
}