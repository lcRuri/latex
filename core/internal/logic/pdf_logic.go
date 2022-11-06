package logic

import (
	"context"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PdfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPdfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PdfLogic {
	return &PdfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PdfLogic) Pdf(req *types.PdfRequest) (resp *types.PdfReply, err error) {
	// todo: add your logic here and delete this line
	//resp.ResponseData.Code = 200
	//resp.ResponseData.Msg = "文件下载成功"
	return
}
