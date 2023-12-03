package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkflowLogic {
	return &GetWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetWorkflowLogic) GetWorkflow(req *types.WorkflowGetRequest) (resp *types.WorkflowRespond, err error) {
	// todo: add your logic here and delete this line

	return
}
