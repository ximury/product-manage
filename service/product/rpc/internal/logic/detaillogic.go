package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"product/service/product/model"

	"product/service/product/rpc/internal/svc"
	"product/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(ctx context.Context, in *product.DetailRequest) (*product.DetailResponse, error) {
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &product.DetailResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  int64(res.Stock),
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil
}
