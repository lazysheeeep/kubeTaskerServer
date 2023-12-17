package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeploymentLogic {
	return &UpdateDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新deployment
func (l *UpdateDeploymentLogic) UpdateDeployment(in *core.UpdateDeploymentReq) (*core.UpdateDeploymentResp, error) {
	// todo: add your logic here and delete this line
	deployment := &Deployment{}
	resp, err := deployment.UpdateDeployment(l, in)
	return resp, err
}
