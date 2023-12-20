package k8snamespace

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamespacesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNamespacesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamespacesLogic {
	return &GetNamespacesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetNamespacesLogic) GetNamespaces(req *types.GetNamespacesReq) (resp *types.GetNamespacesResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.GetNamespaces(l.ctx, &core.GetNamespacesReq{
		FilterName: req.FilterName,
		Page:       req.Page,
		Limit:      req.Limit,
	})
	return &types.GetNamespacesResp{
		Msg: result.Msg,
		Data: types.GetNamespacesData{
			Items: result.Data.Items,
			Total: result.Data.Total,
		},
	}, nil
}
