package k8spod

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodNumPerNpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodNumPerNpLogic {
	return &GetPodNumPerNpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodNumPerNpLogic) GetPodNumPerNp(in *core.GetPodNumPerNpReq) (*core.GetPodNumPerNpResp, error) {
	// todo: add your logic here and delete this line
	pod := Pod{}
	resp, err := pod.GetPodNumPerNp(l, in)
	return resp, err
}
