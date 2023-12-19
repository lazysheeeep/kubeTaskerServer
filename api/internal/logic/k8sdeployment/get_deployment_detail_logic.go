package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeploymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentDetailLogic {
	return &GetDeploymentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetDeploymentDetailLogic) GetDeploymentDetail(req *types.GetDeploymentDetailReq) (resp *types.GetDeploymentDetailResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.GetDeploymentDetail(l.ctx, &core.GetDeploymentDetailReq{
		DeploymentName: req.DeploymentName,
		Namespace:      req.Namespace,
	})
	return &types.GetDeploymentDetailResp{
		Msg:  result.Msg,
		Data: result.Data,
	}, nil
}
