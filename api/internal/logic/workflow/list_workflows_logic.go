package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWorkflowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListWorkflowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWorkflowsLogic {
	return &ListWorkflowsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ListWorkflowsLogic) ListWorkflows(req *types.WorkflowListRequest) (resp *types.WorkflowListRespond, err error) {
	// todo: add your logic here and delete this line
	workflowList, err := l.svcCtx.CoreRpc.ListWorkflows(l.ctx, &core.WorkflowListRequest{
		Namespace:   req.Namespace,
		ListOptions: req.ListOptions,
		Fields:      req.Fields,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowListRespond{
		WorkflowList: workflowList.WorkflowList,
	}, nil
}
