package logic

import (
	"context"
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

func (l *LatexGenerateLogic) LatexGenerate(req *types.LatexGenerateRequest, userIdentity string) (resp *types.LatexGenerateReply, err error) {
	// todo: add your logic here and delete this line
	pdf := &models.Pdf{
		T1:  req.T1,
		T2:  req.T2,
		T3:  req.T3,
		T4:  req.T4,
		T5:  req.T5,
		T6:  req.T6,
		T7:  req.T7,
		T8:  req.T8,
		T9:  req.T9,
		T10: req.T10,
		T11: req.T11,
		T12: req.T12,
		T13: req.T13,
		T14: req.T14,
		T15: req.T15,
		T16: req.T16,
		T17: req.T17,
		T18: req.T18,
		T19: req.T19,
		T20: req.T20,
	}

	filename := pdf.T3 + pdf.T5
	filepath := "./file/njupt" + filename + ".tex"
	helper.PdfGenerate(pdf, filename, filepath)
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
