package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	v1 "k8s.io/api/core/v1"

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
	result, _ := l.svcCtx.CoreRpc.GetServices(l.ctx, &core.GetServicesReq{
		FilterName: req.FilterName,
		Namespace:  req.Namespace,
		Limit:      req.Limit,
		Page:       req.Page,
	})
	items := make([]*v1.Service, 0)
	items = append(items, result.Data.Items...)

	return &types.GetServicesResp{
		Msg: result.Msg,
		Data: &types.GetServicesData{
			Items: items,
			Total: result.Data.Total,
		},
	}, nil
}
