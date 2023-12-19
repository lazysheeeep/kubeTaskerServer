package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigMapLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigMapLogic {
	return &DeleteConfigMapLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteConfigMapLogic) DeleteConfigMap(in *core.DeleteConfigMapReq) (*core.DeleteConfigMapResp, error) {
	// todo: add your logic here and delete this line
	configMap := &ConfigMap{}
	resp, _ := configMap.DeleteConfigMap(l, in)
	return resp, nil

}
