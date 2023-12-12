package k8sPod

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodLogLogic {
	return &GetPodLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodLogLogic) GetPodLog(in *core.GetPodLogReq) (*core.GetPodLogResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetPodLogResp{}, nil
}
