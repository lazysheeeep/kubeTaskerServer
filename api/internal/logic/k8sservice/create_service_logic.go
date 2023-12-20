package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateServiceLogic) CreateService(req *types.CreateServiceReq) (resp *types.CreateServiceResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.CreateService(l.ctx, &core.CreateServiceReq{
		Name:          req.Name,
		Namespace:     req.Namespace,
		Type:          req.Type,
		ContainerPort: req.ContainerPort,
		Port:          req.Port,
		NodePort:      req.NodePort,
		Label:         req.Label,
	})
	return &types.CreateServiceResp{
		Msg: result.Msg,
	}, nil
}
