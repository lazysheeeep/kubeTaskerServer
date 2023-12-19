package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新configmap
func NewUpdateConfigMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigMapLogic {
	return &UpdateConfigMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateConfigMapLogic) UpdateConfigMap(req *types.UpdateConfigMapReq) (resp *types.UpdateConfigMapResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.UpdateConfigMap(l.ctx, &core.UpdateConfigMapReq{
		Namespace: req.Namespace,
		Content:   req.Content,
	})
	return &types.UpdateConfigMapResp{
		Msg: result.Msg,
	}, err
}
