package k8snamespace

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNamespaceLogic {
	return &DeleteNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteNamespaceLogic) DeleteNamespace(req *types.DeleteNamespaceReq) (resp *types.DeleteNamespaceResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.DeleteNamespace(l.ctx, &core.DeleteNamespaceReq{
		NamespaceName: req.NamespaceName,
	})
	return &types.DeleteNamespaceResp{
		Msg:  result.Msg,
		Data: result.Data,
	}, err
}
