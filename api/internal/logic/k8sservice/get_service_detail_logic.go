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
	result, _ := l.svcCtx.CoreRpc.GetServiceDetail(l.ctx, &core.GetServiceDetailReq{
		ServiceName: req.ServiceName,
		Namespace:   req.Namespace,
	})
	return &types.GetServiceDetailResp{
		Msg:  result.Msg,
		Data: result.Data,
	}, nil
}
