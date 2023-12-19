package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceLogic {
	return &DeleteServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteServiceLogic) DeleteService(in *core.DeleteServiceReq) (*core.DeleteServiceResp, error) {
	// todo: add your logic here and delete this line
	service := &Service{}
	resp, _ := service.DeleteService(l, in)
	return resp, nil
}
