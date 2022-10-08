package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"product/service/product/model"

	"product/service/product/rpc/internal/svc"
	"product/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(ctx context.Context, in *product.RemoveRequest) (*product.RemoveResponse, error) {
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.ProductModel.Delete(ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.RemoveResponse{}, nil
}
