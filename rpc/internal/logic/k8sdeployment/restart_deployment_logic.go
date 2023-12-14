package k8sdeployment

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRestartDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartDeploymentLogic {
	return &RestartDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RestartDeploymentLogic) RestartDeployment(in *core.RestartDeploymentReq) (*core.RestartDeploymentResp, error) {
	// todo: add your logic here and delete this line

	return &core.RestartDeploymentResp{}, nil
}
