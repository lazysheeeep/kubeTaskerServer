package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigMapLogic {
	return &DeleteConfigMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteConfigMapLogic) DeleteConfigMap(req *types.DeleteConfigMapReq) (resp *types.DeleteConfigMapResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CoreRpc.DeleteConfigMap(l.ctx, &core.DeleteConfigMapReq{
		ConfigMapName: req.ConfigMapName,
		Namespace:     req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return nil, err
}
