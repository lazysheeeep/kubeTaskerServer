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

func (l *CreateDeploymentLogic) CreateDeployment(in *core.CreateDeploymentReq) (*core.CreateDeploymentResp, error) {
	// todo: add your logic here and delete this line

	return &core.CreateDeploymentResp{}, nil
}
