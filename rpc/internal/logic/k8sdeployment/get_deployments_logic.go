package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeploymentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentsLogic {
	return &GetDeploymentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// K8sDeployment management
func (l *GetDeploymentsLogic) GetDeployments(in *core.GetDeploymentsReq) (*core.GetDeploymentsResp, error) {
	// todo: add your logic here and delete this line
	deployment := &Deployment{}
	resp, err := deployment.GetDeployments(l, in)
	return resp, err
}
