package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"product/service/product/model"

	"product/service/product/rpc/internal/svc"
	"product/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(ctx context.Context, in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Name != "" {
		res.Name = in.Name
	}
	if in.Desc != "" {
		res.Desc = in.Desc
	}
	if in.Stock != 0 {
		res.Stock = uint64(in.Stock)
	}
	if in.Amount != 0 {
		res.Amount = uint64(in.Amount)
	}
	if in.Status != 0 {
		res.Status = uint64(in.Status)
	}

	err = l.svcCtx.ProductModel.Update(ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.UpdateResponse{}, nil
}
