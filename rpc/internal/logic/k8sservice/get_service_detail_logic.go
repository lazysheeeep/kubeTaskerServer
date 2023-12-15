package k8sservice

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceDetailLogic {
	return &GetServiceDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceDetailLogic) GetServiceDetail(in *core.GetServiceDetailReq) (*core.GetServiceDetailResp, error) {
	// todo: add your logic here and delete this line
	service, err := l.svcCtx.K8s.CoreV1().Services(in.Namespace).Get(context.TODO(), in.ServiceName, metav1.GetOptions{})
	if err != nil {
		l.Error(errors.New("获取Service详情失败, " + err.Error()))
		return nil, errors.New("获取Service详情失败, " + err.Error())
	}

	return &core.GetServiceDetailResp{
		Service: service,
	}, nil
}
