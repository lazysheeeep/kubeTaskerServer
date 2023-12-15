package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServicesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServicesLogic {
	return &GetServicesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetServicesLogic) GetServices(req *types.GetServicesReq) (resp *types.GetServicesResp, err error) {
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
