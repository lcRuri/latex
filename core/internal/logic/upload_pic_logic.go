package logic

import (
	"context"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPicLogic {
	return &UploadPicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPicLogic) UploadPic(req *types.PicRequest) (resp *types.PicReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.PicReply)
	resp.ResponseData.Code = 200
	resp.ResponseData.Msg = "上传成功"
	return
}
