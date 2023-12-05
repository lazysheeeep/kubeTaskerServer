package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWorkflowLogic {
	return &CreateWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateWorkflowLogic) CreateWorkflow(req *types.WorkflowCreateRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line
	workflow, err := l.svcCtx.CoreRpc.CreateWorkflow(l.ctx, &core.WorkflowCreateRequest{
		Namespace:     req.Namespace,
		Workflow:      req.Workflow,
		InstanceID:    req.InstanceID,
		ServerDryRun:  req.ServerDryRun,
		CreateOptions: req.CreateOptions,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowRespond{
		Workflow: workflow.Workflow,
	}, nil
}
