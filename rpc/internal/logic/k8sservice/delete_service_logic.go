package k8sservice

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceLogic {
	return &DeleteServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteServiceLogic) DeleteService(in *core.DeleteServiceReq) (*core.DeleteServiceResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.K8s.CoreV1().Services(in.Namespace).Delete(context.TODO(), in.ServiceName, metav1.DeleteOptions{})
	if err != nil {
		l.Error(errors.New("删除Service失败, " + err.Error()))
		return nil, errors.New("删除Service失败, " + err.Error())
	}

	return nil, nil
}
