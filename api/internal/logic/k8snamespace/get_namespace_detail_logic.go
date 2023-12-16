package k8snamespace

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamespaceDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNamespaceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamespaceDetailLogic {
	return &GetNamespaceDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetNamespaceDetailLogic) GetNamespaceDetail(req *types.GetNamespaceDetailReq) (resp *types.GetNamespaceDetailResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetNamespaceDetail(l.ctx, &core.GetNamespaceDetailReq{NamespaceName: req.NamespaceName})
	if err != nil {
		return nil, err
	}
	return &types.GetNamespaceDetailResp{
		Msg: result.Msg,
		Data: types.GetNamespaceDetailData{
			Namespace: result.Data.Namespace,
		},
	}, nil
}
