package k8spod

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodDetailLogic {
	return &GetPodDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodDetailLogic) GetPodDetail(in *core.GetPodDetailReq) (*core.GetPodDetailResp, error) {
	// todo: add your logic here and delete this line
	pod, err := l.svcCtx.K8s.CoreV1().Pods(in.Namespace).Get(context.TODO(), in.PodName, metav1.GetOptions{})
	if err != nil {
		l.Error("获取Pod详情失败," + err.Error())
		return nil, errors.New("获取Pod详情失败," + err.Error())
	}

	return &core.GetPodDetailResp{
		Pod: pod,
	}, nil
}
