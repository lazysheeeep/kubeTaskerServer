package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodNumPerNpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodNumPerNpLogic {
	return &GetPodNumPerNpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPodNumPerNpLogic) GetPodNumPerNp(req *types.GetPodNumPerNpReq) (resp *types.GetPodNumPerNpResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetPodNumPerNp(l.ctx, &core.GetPodNumPerNpReq{})
	if err != nil {
		return nil, err
	}
	podsNps := make([]*types.PodsNp, len(result.PodsNps))
	for _, v := range result.PodsNps {
		podsNps = append(podsNps, &types.PodsNp{
			Namespace: v.Namespace,
			PodNum:    v.PodNum,
		})
	}
	return &types.GetPodNumPerNpResp{
		PodsNps: podsNps,
	}, err
}
