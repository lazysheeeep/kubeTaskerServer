package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuspendWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuspendWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuspendWorkflowLogic {
	return &SuspendWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *SuspendWorkflowLogic) SuspendWorkflow(req *types.WorkflowSuspendRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line
	workflow, err := l.svcCtx.CoreRpc.SuspendWorkflow(l.ctx, &core.WorkflowSuspendRequest{
		Name:      req.Name,
		Namespace: req.Namespace,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowRespond{
		Workflow: workflow.Workflow,
	}, nil
}
