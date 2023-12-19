package k8sPod

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePodLogic {
	return &DeletePodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePodLogic) DeletePod(in *core.DeletePodReq) (*core.DeletePodResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.K8s.CoreV1().Pods(in.Namespace).Delete(context.TODO(), in.PodName, metav1.DeleteOptions{})
	if err != nil {
		l.Error("删除Pod详情失败," + err.Error())
		return nil, errors.New("删除Pod详情失败," + err.Error())
	}
	return nil, nil
}
