package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"product/service/product/model"

	"product/service/product/rpc/internal/svc"
	"product/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CreateLogic) Create(ctx context.Context, in *product.CreateRequest) (*product.CreateResponse, error) {
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  uint64(in.Stock),
		Amount: uint64(in.Amount),
		Status: uint64(in.Status),
	}

	res, err := l.svcCtx.ProductModel.Insert(ctx, &newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	lastInsertId, err := res.LastInsertId()
	newProduct.Id = uint64(lastInsertId)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.CreateResponse{
		Id: int64(newProduct.Id),
	}, nil
}
