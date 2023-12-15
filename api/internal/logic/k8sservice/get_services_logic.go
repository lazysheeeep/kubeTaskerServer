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
	result, err := l.svcCtx.CoreRpc.GetServices(l.ctx, &core.GetServicesReq{
		FilterName: req.FilterName,
		Namespace:  req.Namespace,
		Limit:      req.Limit,
		Page:       req.Page,
	})
	items := make([]*v1.Service, len(result.Items))
	for _, v := range result.Items {
		items = append(items, v)
	}
	if err != nil {
		return nil, err
	}
	return &types.GetServicesResp{
		Items: items,
		Total: result.Total,
	}, err
}
