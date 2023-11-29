package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetryWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetryWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetryWorkflowLogic {
	return &RetryWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RetryWorkflowLogic) RetryWorkflow(in *core.WorkflowRetryRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line

	return &core.WorkflowRespond{}, nil
}
