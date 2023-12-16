package k8snode

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodesLogic {
	return &GetNodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetNodesLogic) GetNodes(req *types.GetNodesReq) (resp *types.GetNodesResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetNodes(l.ctx, &core.GetNodesReq{
		FilterName: req.FilterName,
		Page:       req.Page,
		Limit:      req.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetNodesResp{
		Msg: result.Msg,
		Data: types.GetNodesData{
			Items: result.Data.Items,
			Total: result.Data.Total,
		},
	}, nil
}
