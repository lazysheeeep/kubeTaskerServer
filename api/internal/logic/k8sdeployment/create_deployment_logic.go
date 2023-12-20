package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeploymentLogic {
	return &CreateDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateDeploymentLogic) CreateDeployment(req *types.CreateDeploymentReq) (resp *types.CreateDeploymentResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.CreateDeployment(l.ctx, &core.CreateDeploymentReq{
		Name:          req.Name,
		Namespace:     req.Namespace,
		Replicas:      req.Replicas,
		Image:         req.Image,
		Label:         req.Label,
		Cpu:           req.Cpu,
		Memory:        req.Memory,
		ContainerPort: req.ContainerPort,
		HealthCheck:   req.HealthCheck,
		HealthPath:    req.HealthPath,
	})
	return &types.CreateDeploymentResp{
		Msg: result.Msg,
	}, nil
}
