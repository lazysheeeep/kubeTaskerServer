package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitWorkflowLogic {
	return &SubmitWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SubmitWorkflowLogic) SubmitWorkflow(req *types.WorkflowSubmitRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line
	workflow, err := l.svcCtx.CoreRpc.SubmitWorkflow(l.ctx, &core.WorkflowSubmitRequest{
		Namespace:     req.Namespace,
		ResourceKind:  req.ResourceKind,
		ResourceName:  req.ResourceName,
		SubmitOptions: req.SubmitOptions,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowRespond{
		Workflow: workflow.Workflow,
	}, nil
}
