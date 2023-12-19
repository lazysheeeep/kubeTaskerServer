package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	v1 "k8s.io/api/apps/v1"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeploymentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentsLogic {
	return &GetDeploymentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetDeploymentsLogic) GetDeployments(req *types.GetDeploymentsReq) (resp *types.GetDeploymentsResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.GetDeployments(l.ctx, &core.GetDeploymentsReq{
		FilterName: req.FilterName,
		Namespace:  req.Namespace,
		Limit:      req.Limit,
		Page:       req.Page,
	})
	if result.Data == nil {
		return &types.GetDeploymentsResp{
			Msg:  result.Meg,
			Data: nil,
		}, nil
	}
	items := make([]*v1.Deployment, 0)
	for _, v := range result.Data.Items {
		items = append(items, v)
	}
	return &types.GetDeploymentsResp{
		Msg: result.Meg,
		Data: &types.GetDeploymentsData{
			Items: items,
			Total: result.Data.Total,
		},
	}, nil
}
