package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LintWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLintWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LintWorkflowLogic {
	return &LintWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *LintWorkflowLogic) LintWorkflow(req *types.WorkflowLintRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line

	return
}
