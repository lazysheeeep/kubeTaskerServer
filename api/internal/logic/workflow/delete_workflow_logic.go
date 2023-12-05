package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWorkflowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteWorkflowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWorkflowLogic {
	return &DeleteWorkflowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteWorkflowLogic) DeleteWorkflow(req *types.WorkflowDeleteRequest) (resp *types.WorkflowDeleteResponse, err error) {
	_, err = l.svcCtx.CoreRpc.DeleteWorkflow(l.ctx, &core.WorkflowDeleteRequest{
		Name:          req.Name,
		Namespace:     req.Namespace,
		DeleteOptions: req.DeleteOptions,
		Force:         req.Force,
	})
	if err != nil {
		return nil, err
	}
	return &types.WorkflowDeleteResponse{}, nil
}
