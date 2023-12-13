package k8sconfigmap

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigMapLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigMapLogic {
	return &DeleteConfigMapLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteConfigMapLogic) DeleteConfigMap(in *core.DeleteConfigMapReq) (*core.DeleteConfigMapResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.K8s.CoreV1().ConfigMaps(in.Namespace).Delete(context.TODO(), in.ConfigMapName, metav1.DeleteOptions{})
	if err != nil {
		l.Error(errors.New("删除ConfigMap失败, " + err.Error()))
		return nil, errors.New("删除ConfigMap失败, " + err.Error())
	}
	return nil, nil
}
