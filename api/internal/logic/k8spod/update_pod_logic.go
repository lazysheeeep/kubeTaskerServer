package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePodLogic {
	return &UpdatePodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdatePodLogic) UpdatePod(req *types.UpdatePodReq) (resp *types.UpdatePodResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.UpdatePod(l.ctx, &core.UpdatePodReq{
		Namespace: req.Namespace,
		Content:   req.Content,
	})
	return &types.UpdatePodResp{
		Msg: result.Msg,
	}, nil
}
