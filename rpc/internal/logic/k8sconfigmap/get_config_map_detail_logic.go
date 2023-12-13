package k8sconfigmap

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigMapDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigMapDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigMapDetailLogic {
	return &GetConfigMapDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigMapDetailLogic) GetConfigMapDetail(in *core.GetConfigMapDetailReq) (*core.GetConfigMapDetailResp, error) {
	// todo: add your logic here and delete this line
	configMap, err := l.svcCtx.K8s.CoreV1().ConfigMaps(in.Namespace).Get(context.TODO(), in.ConfigMapName, metav1.GetOptions{})
	if err != nil {
		l.Error(errors.New("获取ConfigMap详情失败, " + err.Error()))
		return nil, errors.New("获取ConfigMap详情失败, " + err.Error())
	}
	return &core.GetConfigMapDetailResp{
		ConfigMap: configMap,
	}, nil
}
