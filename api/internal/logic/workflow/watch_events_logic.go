package workflow

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	"io"

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

func (l *WatchEventsLogic) WatchEvents(req *types.WatchEventsRequest) (resp *types.EventList, err error) {
	// todo: add your logic here and delete this line
	streamEvent, err := l.svcCtx.CoreRpc.WatchEvents(l.ctx, &core.WatchEventsRequest{
		Namespace:   req.Namespace,
		ListOptions: req.ListOptions,
	})
	if err != nil {
		return nil, err
	}
	for {
		res, err := streamEvent.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		resp.EventList = append(resp.EventList, types.Event{
			Event: res.Event,
		})
	}
	return resp, nil
}
