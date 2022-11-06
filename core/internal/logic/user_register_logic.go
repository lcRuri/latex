package logic

import (
	"context"
	"core/helper"
	"core/models"
	"log"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegiserRequest) (resp *types.UserRegiserReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserRegiserReply)

	count, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		log.Printf("models.Engine.Where err", err)
		resp.ResponseData.Code = 200
		resp.ResponseData.Msg = "用户已存在"
		return
	}

	if count > 0 {
		l.Logger.Error("count>0")
		resp.ResponseData.Code = 200
		resp.ResponseData.Msg = "用户已存在"
		return
	}

	//用户注册
	if req.Password != req.RePassword {
		resp.ResponseData.Code = 200
		resp.ResponseData.Msg = "前后密码不一致"
		return
	}
	user := &models.UserBasic{
		Identity: helper.UUid(),
		Name:     req.Name,
		Password: req.Password,
	}

	n, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		log.Printf("models.Engine.Insert failed", err)
		return
	}

	if n != 1 {
		log.Printf("models.Engine.Insert failed")
		return
	}

	resp.ResponseData.Code = 200
	resp.ResponseData.Msg = "注册成功"
	return
}
