package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	"io"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WatchWorkflowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWatchWorkflowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WatchWorkflowsLogic {
	return &WatchWorkflowsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *WatchWorkflowsLogic) WatchWorkflows(req *types.WatchWorkflowsRequest) (resp *types.WorkflowWatchEventList, err error) {
	// todo: add your logic here and delete this line
	streamWorkflowWatchEvent, err := l.svcCtx.CoreRpc.WatchWorkflows(l.ctx, &core.WatchWorkflowsRequest{
		Namespace:   req.Namespace,
		ListOptions: req.ListOptions,
		Fields:      req.Fields,
	})
	if err != nil {
		return nil, err
	}
	for {
		res, err := streamWorkflowWatchEvent.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		resp.WorkflowWatchEventList = append(resp.WorkflowWatchEventList, types.WorkflowWatchEvent{
			Type:   res.Type,
			Object: res.Object,
		})
	}
	return resp, nil
}
