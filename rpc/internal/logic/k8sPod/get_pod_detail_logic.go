package k8sPod

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodDetailLogic {
	return &GetPodDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodDetailLogic) GetPodDetail(in *core.GetPodDetailReq) (*core.GetPodDetailResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetPodDetailResp{}, nil
}
