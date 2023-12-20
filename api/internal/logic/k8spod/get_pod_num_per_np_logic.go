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
	result, _ := l.svcCtx.CoreRpc.GetPodNumPerNp(l.ctx, &core.GetPodNumPerNpReq{})
	getPodNumPerNpData := make([]*types.GetPodNumPerNpData, 0)
	for _, v := range result.Data {
		getPodNumPerNpData = append(getPodNumPerNpData, &types.GetPodNumPerNpData{
			Namespace: v.Namespace,
			PodNum:    v.PodNum,
		})
	}
	return &types.GetPodNumPerNpResp{
		Msg:  result.Msg,
		Data: getPodNumPerNpData,
	}, nil
}
