package k8sdeployment

import (
	"context"
	"encoding/json"
	"errors"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeploymentLogic {
	return &UpdateDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新deployment
func (l *UpdateDeploymentLogic) UpdateDeployment(in *core.UpdateDeploymentReq) (*core.UpdateDeploymentResp, error) {
	// todo: add your logic here and delete this line
	var deploy = &v1.Deployment{}

	err := json.Unmarshal([]byte(in.Content), deploy)
	if err != nil {
		l.Error(errors.New("反序列化失败, " + err.Error()))
		return nil, errors.New("反序列化失败, " + err.Error())
	}

	_, err = l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		l.Error(errors.New("更新Deployment失败, " + err.Error()))
		return nil, errors.New("更新Deployment失败, " + err.Error())
	}
	return nil, nil
}
