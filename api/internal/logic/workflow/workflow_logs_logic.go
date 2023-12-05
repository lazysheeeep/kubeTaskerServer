package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	"io"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkflowLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkflowLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkflowLogsLogic {
	return &WorkflowLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *WorkflowLogsLogic) WorkflowLogs(req *types.WorkflowLogRequest) (resp *types.LogEntryList, err error) {
	// todo: add your logic here and delete this line
	streamLogEntry, err := l.svcCtx.CoreRpc.WorkflowLogs(l.ctx, &core.WorkflowLogRequest{
		Name:       req.Name,
		Namespace:  req.Namespace,
		PodName:    req.PodName,
		LogOptions: req.LogOptions,
		Grep:       req.Grep,
		Selector:   req.Selector,
	})
	if err != nil {
		return nil, err
	}
	for {
		res, err := streamLogEntry.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		resp.LogEntryList = append(resp.LogEntryList, types.LogEntry{
			Content: res.Content,
			PodName: res.PodName,
		})
	}
	return resp, nil
}
