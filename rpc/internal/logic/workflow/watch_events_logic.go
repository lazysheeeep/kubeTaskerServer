package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type WatchEventsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWatchEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WatchEventsLogic {
	return &WatchEventsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WatchEventsLogic) WatchEvents(in *core.WatchEventsRequest, stream core.Core_WatchEventsServer) error {
	// todo: add your logic here and delete this line

	return nil
}
