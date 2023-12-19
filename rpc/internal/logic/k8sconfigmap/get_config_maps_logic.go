package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigMapsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigMapsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigMapsLogic {
	return &GetConfigMapsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// K8sConfigMap management
func (l *GetConfigMapsLogic) GetConfigMaps(in *core.GetConfigMapsReq) (*core.GetConfigMapsResp, error) {
	// todo: add your logic here and delete this line
	configMap := &ConfigMap{}
	resp, _ := configMap.GetConfigMaps(l, in)
	return resp, nil
}
