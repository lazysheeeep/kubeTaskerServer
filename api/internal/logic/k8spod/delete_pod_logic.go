package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePodLogic {
	return &DeletePodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeletePodLogic) DeletePod(req *types.DeletePodReq) (resp *types.DeletePodResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.DeletePod(l.ctx, &core.DeletePodReq{
		PodName:   req.PodName,
		Namespace: req.Namespace,
	})
	return &types.DeletePodResp{
		Msg: result.Msg,
	}, err
}
