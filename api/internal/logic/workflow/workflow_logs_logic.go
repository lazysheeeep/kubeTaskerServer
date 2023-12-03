package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkflowLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkflowLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkflowLogsLogic {
	return &WorkflowLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *WorkflowLogsLogic) WorkflowLogs(req *types.WorkflowLogRequest) (resp *types.LogEntry, err error) {
	// todo: add your logic here and delete this line

	return
}
