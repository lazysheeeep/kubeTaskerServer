package k8sdeployment

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeploymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentDetailLogic {
	return &GetDeploymentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeploymentDetailLogic) GetDeploymentDetail(in *core.GetDeploymentDetailReq) (*core.GetDeploymentDetailResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetDeploymentDetailResp{}, nil
}
