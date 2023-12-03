package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitWorkflowLogic {
	return &SubmitWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitWorkflowLogic) SubmitWorkflow(in *core.WorkflowSubmitRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line

	return &core.WorkflowRespond{}, nil
}
