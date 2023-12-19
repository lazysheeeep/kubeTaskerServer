package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScaleDeploymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScaleDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScaleDeploymentLogic {
	return &ScaleDeploymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ScaleDeploymentLogic) ScaleDeployment(req *types.ScaleDeploymentReq) (resp *types.ScaleDeploymentResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.ScaleDeployment(l.ctx, &core.ScaleDeploymentReq{
		DeploymentName: req.DeploymentName,
		Namespace:      req.Namespace,
		ScaleNum:       req.ScaleNum,
	})
	if err != nil {
		return nil, err
	}
	return &types.ScaleDeploymentResp{
		Msg:  result.Msg,
		Data: result.Data,
	}, err
}
