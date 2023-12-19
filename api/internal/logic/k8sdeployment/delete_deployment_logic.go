package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeploymentLogic {
	return &DeleteDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteDeploymentLogic) DeleteDeployment(req *types.DeleteDeploymentReq) (resp *types.DeleteDeploymentResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.DeleteDeployment(l.ctx, &core.DeleteDeploymentReq{
		DeploymentName: req.DeploymentName,
		Namespace:      req.Namespace,
	})
	return &types.DeleteDeploymentResp{
		Msg: result.Msg,
	}, err
}
