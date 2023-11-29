package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodLogsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPodLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PodLogsLogic {
	return &PodLogsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PodLogsLogic) PodLogs(in *core.WorkflowLogRequest, stream core.Core_PodLogsServer) error {
	// todo: add your logic here and delete this line

	return nil
}
