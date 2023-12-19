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

// 获取configmap列表，支持过滤、排序、分页
func (l *GetConfigMapsLogic) GetConfigMaps(req *types.GetConfigMapsReq) (resp *types.GetConfigMapsResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.CoreRpc.GetConfigMaps(l.ctx, &core.GetConfigMapsReq{
		FilterName: req.FilterName,
		Namespace:  req.Namespace,
		Limit:      req.Limit,
		Page:       req.Page,
	})
	if result.Data == nil {
		resp = &types.GetConfigMapsResp{
			Msg:  result.Msg,
			Data: nil,
		}
	} else {
		resp = &types.GetConfigMapsResp{
			Msg: result.Msg,
			Data: &types.GetConfigMapsData{
				Items: result.Data.Items,
				Total: result.Data.Total,
			},
		}
	}
	return resp, nil
}
