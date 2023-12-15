package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodDetailLogic {
	return &GetPodDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPodDetailLogic) GetPodDetail(req *types.GetPodDetailReq) (resp *types.GetPodDetailResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetPodDetail(l.ctx, &core.GetPodDetailReq{
		PodName:   req.PodName,
		Namespace: req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetPodDetailResp{
		Pod: result.Pod,
	}, err
}
