package k8sdeployment

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeploymentLogic {
	return &DeleteDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除deployment
func (l *DeleteDeploymentLogic) DeleteDeployment(in *core.DeleteDeploymentReq) (*core.DeleteDeploymentResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).Delete(context.TODO(), in.DeploymentName, metav1.DeleteOptions{})
	if err != nil {
		l.Error(errors.New("删除Deployment失败, " + err.Error()))
		return nil, errors.New("删除Deployment失败, " + err.Error())
	}

	return nil, nil
}
