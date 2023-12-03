package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetWorkflowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetWorkflowLogic {
	return &SetWorkflowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetWorkflowLogic) SetWorkflow(in *core.WorkflowSetRequest) (*core.WorkflowRespond, error) {
	// todo: add your logic here and delete this line

	return &core.WorkflowRespond{}, nil
}
