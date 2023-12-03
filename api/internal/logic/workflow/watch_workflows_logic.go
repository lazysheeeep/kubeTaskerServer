package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WatchWorkflowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWatchWorkflowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WatchWorkflowsLogic {
	return &WatchWorkflowsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *WatchWorkflowsLogic) WatchWorkflows(req *types.WatchWorkflowsRequest) (resp *types.WorkflowWatchEvent, err error) {
	// todo: add your logic here and delete this line

	return
}
