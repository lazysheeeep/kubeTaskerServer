package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PodLogsLogic {
	return &PodLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *PodLogsLogic) PodLogs(req *types.WorkflowLogRequest) (resp *types.LogEntry, err error) {
	// todo: add your logic here and delete this line

	return
}
