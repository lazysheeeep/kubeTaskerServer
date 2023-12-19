package k8sPod

import (
	"context"
	"encoding/json"
	"errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePodLogic {
	return &UpdatePodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePodLogic) UpdatePod(in *core.UpdatePodReq) (*core.UpdatePodResp, error) {
	// todo: add your logic here and delete this line
	var pod = &corev1.Pod{}
	// 反序列化为Pod对象
	err := json.Unmarshal([]byte(in.Content), pod)
	if err != nil {
		l.Error("反序列化失败," + err.Error())
		return nil, errors.New("反序列化失败," + err.Error())
	}
	// 更新pod
	_, err = l.svcCtx.K8s.CoreV1().Pods(in.Namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		l.Error("更新Pod失败," + err.Error())
		return nil, errors.New("更新Pod失败," + err.Error())
	}
	return nil, nil
}
