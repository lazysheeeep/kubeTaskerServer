package k8sservice

import (
	"context"
	"encoding/json"
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceLogic {
	return &UpdateServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateServiceLogic) UpdateService(in *core.UpdateServiceReq) (*core.UpdateServiceResp, error) {
	// todo: add your logic here and delete this line
	var service = &v1.Service{}

	err := json.Unmarshal([]byte(in.Content), service)
	if err != nil {
		l.Error(errors.New("反序列化失败, " + err.Error()))
		return nil, errors.New("反序列化失败, " + err.Error())
	}

	_, err = l.svcCtx.K8s.CoreV1().Services(in.Namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		l.Error(errors.New("更新service失败, " + err.Error()))
		return nil, errors.New("更新service失败, " + err.Error())
	}
	return nil, nil
}
