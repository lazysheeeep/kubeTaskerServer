package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStopWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopWorkflowLogic {
	return &StopWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StopWorkflowLogic) StopWorkflow(in *core.WorkflowStopRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line

	return &core.WorkflowRespond{}, nil
}
