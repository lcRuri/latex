package logic

import (
	"context"
	"core/define"
	"core/helper"
	"core/internal/svc"
	"core/internal/types"
	"core/models"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LatexGenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLatexGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LatexGenerateLogic {
	return &LatexGenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LatexGenerateLogic) LatexGenerate(req *types.LatexGenerateRequest, userIdentity, content string) (resp *types.LatexGenerateReply, err error) {
	// todo: add your logic here and delete this line

	//filename := pdf.Title + pdf.GroupLeaderName
	filename := helper.UUid()
	filepath := define.BaseUploadPath + "/" + filename + ".tex"
	helper.PdfGenerate(filepath, filename, content)
	if err != nil {
		l.Logger.Error(err)
		return
	}
	rp := new(models.RepositoryPool)
	//查询新生成的文件的hash值是否在数据库已经存在
	hash := helper.Hash(filepath)
	//hash := helper.UUid()
	cnt, err := l.svcCtx.Engine.Where("hash = ?", hash).Count(rp)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	if cnt > 0 {
		return nil, errors.New("该名称已经存在")
	}

	//插入文件数据
	insert := &models.RepositoryPool{
		Identity: userIdentity,
		Hash:     hash,
		Name:     "njupt" + filename,
		Path:     filepath,
	}

	_, err = l.svcCtx.Engine.Insert(insert)
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = new(types.LatexGenerateReply)
	resp.ResponseData.Code = 200
	resp.ResponseData.Msg = "pdf文档生成成功"
	resp.Filename = "njupt" + filename

	return
}
