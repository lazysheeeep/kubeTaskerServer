package k8snode

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeDetailLogic {
	return &GetNodeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetNodeDetailLogic) GetNodeDetail(req *types.GetNodeDetailReq) (resp *types.GetNodeDetailResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetNodeDetail(l.ctx, &core.GetNodeDetailReq{
		NodeName: req.NodeName,
	})
	return &types.GetNodeDetailResp{
		Msg: result.Msg,
		Data: types.GetNodeDetailData{
			Node: result.Data.Node,
		},
	}, err
}
