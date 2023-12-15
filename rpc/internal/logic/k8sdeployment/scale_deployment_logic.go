package k8sdeployment

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScaleDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScaleDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScaleDeploymentLogic {
	return &ScaleDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 设置deployment副本数
func (l *ScaleDeploymentLogic) ScaleDeployment(in *core.ScaleDeploymentReq) (*core.ScaleDeploymentResp, error) {
	// todo: add your logic here and delete this line
	//获取autoscalingv1.Scale类型的对象，能点出当前的副本数
	scale, err := l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).GetScale(context.TODO(), in.DeploymentName, metav1.GetOptions{})
	if err != nil {
		l.Error(errors.New("获取Deployment副本数信息失败, " + err.Error()))
		return &core.ScaleDeploymentResp{
			Replica: 0,
		}, errors.New("获取Deployment副本数信息失败, " + err.Error())
	}
	//修改副本数
	scale.Spec.Replicas = int32(in.ScaleNum)
	//更新副本数，传入scale对象
	newScale, err := l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).UpdateScale(context.TODO(), in.DeploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		l.Error(errors.New("更新Deployment副本数信息失败, " + err.Error()))
		return &core.ScaleDeploymentResp{
			Replica: 0,
		}, errors.New("更新Deployment副本数信息失败, " + err.Error())
	}

	return &core.ScaleDeploymentResp{
		Replica: int64(newScale.Spec.Replicas),
	}, nil
}
