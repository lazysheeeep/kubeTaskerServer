package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RetryWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRetryWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetryWorkflowLogic {
	return &RetryWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RetryWorkflowLogic) RetryWorkflow(req *types.WorkflowRetryRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line
	workflow, err := l.svcCtx.CoreRpc.RetryWorkflow(l.ctx, &core.WorkflowRetryRequest{
		Name:              req.Name,
		Namespace:         req.Namespace,
		RestartSuccessful: req.RestartSuccessful,
		NodeFieldSelector: req.NodeFieldSelector,
		Parameters:        req.Parameters,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowRespond{
		Workflow: workflow.Workflow,
	}, nil
}
