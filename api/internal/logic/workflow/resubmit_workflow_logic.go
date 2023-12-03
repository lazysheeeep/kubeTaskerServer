package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResubmitWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResubmitWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResubmitWorkflowLogic {
	return &ResubmitWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ResubmitWorkflowLogic) ResubmitWorkflow(req *types.WorkflowResubmitRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line

	return
}
