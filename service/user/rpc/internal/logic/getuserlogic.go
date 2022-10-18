package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"product/service/user/model"

	"product/service/user/rpc/internal/svc"
	"product/service/user/rpc/pb/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(ctx context.Context, in *userclient.GetUserRequest) (*userclient.GetUserResponse, error) {
	// 查询用户是否存在
	l.Logger.Infof("%d\n\n", in.Id)
	res, err := l.svcCtx.UserModel.FindOne(ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	l.Logger.Infof("%#v\n", res)
	return &userclient.GetUserResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Gender: int64(res.Gender),
		Mobile: res.Mobile,
	}, nil
}
