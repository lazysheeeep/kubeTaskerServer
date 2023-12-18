package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeploymentLogic {
	return &UpdateDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateDeploymentLogic) UpdateDeployment(req *types.UpdateDeploymentReq) (resp *types.UpdateDeploymentResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.UpdateDeployment(l.ctx, &core.UpdateDeploymentReq{
		Namespace: req.Namespace,
		Content:   req.Content,
	})
	return &types.UpdateDeploymentResp{
		Msg: result.Msg,
	}, err
}
