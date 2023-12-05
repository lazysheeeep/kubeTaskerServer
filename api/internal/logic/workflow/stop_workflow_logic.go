package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopWorkflowLogic {
	return &StopWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *StopWorkflowLogic) StopWorkflow(req *types.WorkflowStopRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line
	workflow, err := l.svcCtx.CoreRpc.StopWorkflow(l.ctx, &core.WorkflowStopRequest{
		Name:              req.Name,
		Namespace:         req.Namespace,
		NodeFieldSelector: req.NodeFieldSelector,
		Message:           req.Message,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowRespond{
		Workflow: workflow.Workflow,
	}, nil
}
