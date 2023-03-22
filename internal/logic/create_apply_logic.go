package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/internal/model"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApplyLogic {
	return &CreateApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateApplyLogic) CreateApply(in *pb.CreateApplyReq) (*pb.CreateApplyResp, error) {
	if err := l.svcCtx.ApplyModel.Insert(l.ctx, &model.Apply{
		ApplicantId: in.ApplicantId,
		CommunityId: in.CommunityId,
	}); err != nil {
		return nil, err
	}

	return &pb.CreateApplyResp{}, nil
}
