package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceDetailLogic {
	return &GetServiceDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetServiceDetailLogic) GetServiceDetail(req *types.GetServiceDetailReq) (resp *types.GetServiceDetailResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetConfigMapDetail(l.ctx, &core.GetConfigMapDetailReq{
		ConfigMapName: req.ConfigMapName,
		Namespace:     req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetConfigMapDetailResp{
		ConfigMap: result.ConfigMap,
	}, err
}
