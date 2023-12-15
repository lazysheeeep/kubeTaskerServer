package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestartDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartDeploymentLogic {
	return &RestartDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RestartDeploymentLogic) RestartDeployment(req *types.RestartDeploymentReq) (resp *types.RestartDeploymentResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CoreRpc.RestartDeployment(l.ctx, &core.RestartDeploymentReq{
		DeploymentName: req.DeploymentName,
		Namespace:      req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return nil, err
}
