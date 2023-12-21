package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	v1 "k8s.io/api/core/v1"
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
	if req.Namespace == "default" {
		req.Namespace = ""
	}
	result, _ := l.svcCtx.CoreRpc.GetPods(l.ctx, &core.GetPodsReq{
		FilterName: req.FilterName,
		Namespace:  req.Namespace,
		Limit:      req.Limit,
		Page:       req.Page,
	})
	items := make([]*v1.Pod, 0)
	items = append(items, result.Data.Items...)

	return &types.GetPodsResp{
		Msg: result.Msg,
		Data: &types.GetPodsData{
			Total: result.Data.Total,
			Items: items,
		},
	}, nil
}
