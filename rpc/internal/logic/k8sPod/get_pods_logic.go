package k8sPod

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodsLogic {
	return &GetPodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// K8sPod management
func (l *GetPodsLogic) GetPods(in *core.GetPodsReq) (*core.GetPodsResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetPodsResp{}, nil
}
