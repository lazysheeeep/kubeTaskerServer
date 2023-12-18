package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigMapDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取configmap详情
func NewGetConfigMapDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigMapDetailLogic {
	return &GetConfigMapDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetConfigMapDetailLogic) GetConfigMapDetail(req *types.GetConfigMapDetailReq) (resp *types.GetConfigMapDetailResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetConfigMapDetail(l.ctx, &core.GetConfigMapDetailReq{
		ConfigMapName: req.ConfigMapName,
		Namespace:     req.Namespace,
	})
	return &types.GetConfigMapDetailResp{
		Msg:  result.Msg,
		Data: result.Data,
	}, err
}
