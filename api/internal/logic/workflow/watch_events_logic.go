package workflow

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WatchEventsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWatchEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WatchEventsLogic {
	return &WatchEventsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *WatchEventsLogic) WatchEvents(req *types.WatchEventsRequest) (resp *types.Event, err error) {
	// todo: add your logic here and delete this line

	return
}
