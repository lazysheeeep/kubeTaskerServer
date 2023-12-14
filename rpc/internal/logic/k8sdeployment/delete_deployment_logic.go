package k8sdeployment

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeploymentLogic {
	return &DeleteDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeploymentLogic) DeleteDeployment(in *core.DeleteDeploymentReq) (*core.DeleteDeploymentResp, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteDeploymentResp{}, nil
}
