package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResumeWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResumeWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResumeWorkflowLogic {
	return &ResumeWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ResumeWorkflowLogic) ResumeWorkflow(req *types.WorkflowResumeRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line

	return
}
