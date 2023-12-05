package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetWorkflowLogic {
	return &SetWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SetWorkflowLogic) SetWorkflow(req *types.WorkflowSetRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line
	workflow, err := l.svcCtx.CoreRpc.SetWorkflow(l.ctx, &core.WorkflowSetRequest{
		Name:              req.Name,
		Namespace:         req.Namespace,
		NodeFieldSelector: req.NodeFieldSelector,
		Message:           req.Message,
		Phase:             req.Phase,
		OutputParameters:  req.OutputParameters,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowRespond{
		Workflow: workflow.Workflow,
	}, nil
}
