package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodContainerLogic {
	return &GetPodContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPodContainerLogic) GetPodContainer(req *types.GetPodContainerReq) (resp *types.GetPodContainerResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetPodContainer(l.ctx, &core.GetPodContainerReq{
		PodName:   req.PodName,
		Namespace: req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetPodContainerResp{
		Containers: result.Containers,
	}, err
}
