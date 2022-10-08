package logic

import (
	"context"
	"product/common/cryptx"
	"product/service/user/model"
	"product/service/user/rpc/internal/svc"
	"product/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}

	if err == model.ErrNotFound {

		newUser := model.User{
			Name:     in.Name,
			Gender:   uint64(in.Gender),
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		lastInsertId, err := res.LastInsertId()
		newUser.Id = uint64(lastInsertId)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:     int64(newUser.Id),
			Name:   newUser.Name,
			Gender: int64(newUser.Gender),
			Mobile: newUser.Mobile,
		}, nil

	}

	return nil, status.Error(500, err.Error())
}
