package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceLogic {
	return &UpdateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateServiceLogic) UpdateService(req *types.UpdateServiceReq) (resp *types.UpdateServiceResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CoreRpc.UpdateService(l.ctx, &core.UpdateServiceReq{
		Namespace: req.Namespace,
		Content:   req.Content,
	})
	if err != nil {
		return nil, err
	}
	return nil, err
}
