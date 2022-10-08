// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package server

import (
	"context"

	"product/service/product/rpc/internal/logic"
	"product/service/product/rpc/internal/svc"
	"product/service/product/rpc/pb/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Create(ctx context.Context, in *product.CreateRequest) (*product.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *ProductServer) Update(ctx context.Context, in *product.UpdateRequest) (*product.UpdateResponse, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *ProductServer) Remove(ctx context.Context, in *product.RemoveRequest) (*product.RemoveResponse, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

func (s *ProductServer) Detail(ctx context.Context, in *product.DetailRequest) (*product.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}
