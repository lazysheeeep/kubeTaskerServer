package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigMapLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigMapLogic {
	return &UpdateConfigMapLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigMapLogic) UpdateConfigMap(in *core.UpdateConfigMapReq) (*core.UpdateConfigMapResp, error) {
	// todo: add your logic here and delete this line
	configMap := &ConfigMap{}
	resp, _ := configMap.UpdateConfigMap(l, in)
	return resp, nil
}
