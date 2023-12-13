package k8spod

import (
	"context"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodContainerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodContainerLogic {
	return &GetPodContainerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodContainerLogic) GetPodContainer(in *core.GetPodContainerReq) (*core.GetPodContainerResp, error) {
	// todo: add your logic here and delete this line
	getPodDetailLogic := NewGetPodDetailLogic(l.ctx, l.svcCtx)
	pod, err := getPodDetailLogic.GetPodDetail(&core.GetPodDetailReq{
		PodName:   in.PodName,
		Namespace: in.Namespace,
	})
	if err != nil {
		return nil, err
	}
	containers := make([]string, len(pod.Pod.Spec.Containers))
	for _, container := range pod.Pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return &core.GetPodContainerResp{
		Containers: containers,
	}, nil
}
