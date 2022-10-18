package logic

import (
	"context"
	"product/service/user/rpc/pb/userclient"

	"product/service/user/api/internal/svc"
	"product/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (resp *types.GetUseResponse, err error) {
	res, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.GetUserRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.GetUseResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
