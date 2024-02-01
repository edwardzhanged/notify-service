package logic

import (
	"context"

	"notify-service/internal/svc"
	"notify-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifyByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyByEmailLogic {
	return &NotifyByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifyByEmailLogic) NotifyByEmail(req *types.NotifyByEmailReq) (resp *types.NotifyByEmailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
