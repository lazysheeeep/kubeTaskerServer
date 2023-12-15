package k8sdeployment

import (
	"context"
	"errors"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeploymentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeploymentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeploymentsLogic {
	return &GetDeploymentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// K8sDeployment management
func (l *GetDeploymentsLogic) GetDeployments(in *core.GetDeploymentsReq) (*core.GetDeploymentsResp, error) {
	// todo: add your logic here and delete this line
	//获取deploymentList类型的deployment列表
	deploymentList, err := l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Error(errors.New("获取Deployment列表失败, " + err.Error()))
		return nil, errors.New("获取Deployment列表失败, " + err.Error())
	}
	//将deploymentList中的deployment列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: toCells(deploymentList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: in.FilterName},
			Paginate: &PaginateQuery{
				Limit: int(in.Limit),
				Page:  int(in.Page),
			},
		},
	}

	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()
	println("Deployment total: ", total)

	//将[]DataCell类型的deployment列表转为appsv1.deployment列表
	deployments := fromCells(data.GenericDataList)
	items := make([]*v1.Deployment, 0)
	for _, v := range deployments {
		items = append(items, &v)
	}
	return &core.GetDeploymentsResp{
		Items: items,
		Total: int64(total),
	}, nil
}
