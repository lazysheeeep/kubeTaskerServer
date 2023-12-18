package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodLogLogic {
	return &GetPodLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPodLogLogic) GetPodLog(req *types.GetPodLogReq) (resp *types.GetPodLogResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetPodLog(l.ctx, &core.GetPodLogReq{
		ContainerName: req.ContainerName,
		PodName:       req.PodName,
		Namespace:     req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetPodLogResp{
		Msg:  result.Meg,
		Data: result.Data,
	}, err
}
