package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodNumPerNpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPodNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodNumPerNpLogic {
	return &GetPodNumPerNpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPodNumPerNpLogic) GetPodNumPerNp(req *types.GetPodNumPerNpReq) (resp *types.GetPodNumPerNpResp, err error) {
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
