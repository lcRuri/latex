package logic

import (
	"context"
	"core/define"
	"core/helper"
	"core/models"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.LoginReply)

	//1.从数据库中查询用户是否存在
	user := new(models.UserBasic)
	ok, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, req.Password).Get(user)
	if err != nil {
		l.Logger.Errorf("models.Engine.Where err:%v", err)
		resp.ResponseData.Code = 200
		resp.ResponseData.Msg = "查询出错"
		return
	}
	if !ok {
		resp.ResponseData.Code = 200
		resp.ResponseData.Msg = "用户名或密码错误"
		return
	}

	//2.生成Token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpireTime)
	if err != nil {
		l.Logger.Errorf("helper.GenerateToken err:", err)
		return nil, err
	}

	//3.返回信息
	resp.ResponseData.Code = 200
	resp.ResponseData.Msg = "success"
	resp.Token = token

	return
}
