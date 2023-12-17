package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeploymentLogic {
	return &CreateDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建deployment,接收DeployCreate对象
func (l *CreateDeploymentLogic) CreateDeployment(in *core.CreateDeploymentReq) (*core.CreateDeploymentResp, error) {
	deployment := &Deployment{}
	resp, err := deployment.CreateDeployment(l, in)
	return resp, err
}
