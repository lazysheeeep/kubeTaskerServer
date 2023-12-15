package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodsLogic {
	return &GetPodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPodsLogic) GetPods(req *types.GetPodsReq) (resp *types.GetPodsResp, err error) {
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
