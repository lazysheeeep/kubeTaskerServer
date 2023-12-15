package k8sdeployment

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeployNumPerNpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeployNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeployNumPerNpLogic {
	return &GetDeployNumPerNpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeployNumPerNpLogic) GetDeployNumPerNp(in *core.GetDeployNumPerNpReq) (*core.GetDeployNumPerNpResp, error) {
	// todo: add your logic here and delete this line
	namespaceList, err := l.svcCtx.K8s.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	deploysNps := make([]*core.DeploysNp, 0)
	for _, namespace := range namespaceList.Items {
		deploymentList, err := l.svcCtx.K8s.AppsV1().Deployments(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		deploysNp := &core.DeploysNp{
			Namespace: namespace.Name,
			DeployNum: int64(len(deploymentList.Items)),
		}

		deploysNps = append(deploysNps, deploysNp)
	}
	return &core.GetDeployNumPerNpResp{
		DeploysNps: deploysNps,
	}, nil
}
