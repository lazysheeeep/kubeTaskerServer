package k8sdeployment

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeploymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentDetailLogic {
	return &GetDeploymentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取deployment详情
func (l *GetDeploymentDetailLogic) GetDeploymentDetail(in *core.GetDeploymentDetailReq) (*core.GetDeploymentDetailResp, error) {
	// todo: add your logic here and delete this line
	deployment, err := l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).Get(context.TODO(), in.DeploymentName, metav1.GetOptions{})
	if err != nil {
		l.Error(errors.New("获取Deployment详情失败, " + err.Error()))
		return nil, errors.New("获取Deployment详情失败, " + err.Error())
	}
	return &core.GetDeploymentDetailResp{
		Deployment: deployment,
	}, nil
}
