package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigMapDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigMapDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigMapDetailLogic {
	return &GetConfigMapDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigMapDetailLogic) GetConfigMapDetail(in *core.GetConfigMapDetailReq) (*core.GetConfigMapDetailResp, error) {
	// todo: add your logic here and delete this line
	configMap := &ConfigMap{}
	resp, _ := configMap.GetConfigMapDetail(l, in)
	return resp, nil
}
