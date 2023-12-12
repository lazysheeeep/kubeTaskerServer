package k8sPod

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePodLogic {
	return &UpdatePodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePodLogic) UpdatePod(in *core.UpdatePodReq) (*core.UpdatePodResp, error) {
	// todo: add your logic here and delete this line

	return &core.UpdatePodResp{}, nil
}
