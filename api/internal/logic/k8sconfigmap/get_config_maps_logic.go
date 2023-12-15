package k8sconfigmap

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigMapsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigMapsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigMapsLogic {
	return &GetConfigMapsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetConfigMapsLogic) GetConfigMaps(req *types.GetConfigMapsReq) (resp *types.GetConfigMapsResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetConfigMaps(l.ctx, &core.GetConfigMapsReq{
		FilterName: req.FilterName,
		Namespace:  req.Namespace,
		Limit:      req.Limit,
		Page:       req.Page,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetConfigMapsResp{
		Items: result.Items,
		Total: resp.Total,
	}, err
}
