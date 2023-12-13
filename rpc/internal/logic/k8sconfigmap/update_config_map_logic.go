package k8sconfigmap

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

type UpdateConfigMapLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigMapLogic {
	return &UpdateConfigMapLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateConfigMapLogic) UpdateConfigMap(in *core.UpdateConfigMapReq) (*core.UpdateConfigMapResp, error) {
	// todo: add your logic here and delete this line
	var configMap = &v1.ConfigMap{}

	err := json.Unmarshal([]byte(in.Content), configMap)
	if err != nil {
		l.Error(errors.New("反序列化失败, " + err.Error()))
		return nil, errors.New("反序列化失败, " + err.Error())
	}

	_, err = l.svcCtx.K8s.CoreV1().ConfigMaps(in.Namespace).Update(context.TODO(), configMap, metav1.UpdateOptions{})
	if err != nil {
		l.Error(errors.New("更新ConfigMap失败, " + err.Error()))
		return nil, errors.New("更新ConfigMap失败, " + err.Error())
	}
	return nil, nil
}
