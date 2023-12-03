package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TerminateWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTerminateWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TerminateWorkflowLogic {
	return &TerminateWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *TerminateWorkflowLogic) TerminateWorkflow(req *types.WorkflowTerminateRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line

	return
}
