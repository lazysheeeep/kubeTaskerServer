package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeployNumPerNpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeployNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeployNumPerNpLogic {
	return &GetDeployNumPerNpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeployNumPerNpLogic) GetDeployNumPerNp(in *core.GetDeployNumPerNpReq) (*core.GetDeployNumPerNpResp, error) {
	// todo: add your logic here and delete this line
	deployment := &Deployment{}
	resp, err := deployment.GetDeployNumPerNp(l, in)
	return resp, err
}
