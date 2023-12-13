package k8sconfigmap

import (
	"context"
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigMapsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigMapsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigMapsLogic {
	return &GetConfigMapsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// K8sConfigMap management
func (l *GetConfigMapsLogic) GetConfigMaps(in *core.GetConfigMapsReq) (*core.GetConfigMapsResp, error) {
	// todo: add your logic here and delete this line
	//获取configMapList类型的configMap列表
	configMapList, err := l.svcCtx.K8s.CoreV1().ConfigMaps(in.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Error(errors.New("获取ConfigMap列表失败, " + err.Error()))
		return nil, errors.New("获取ConfigMap列表失败, " + err.Error())
	}
	//将configMapList中的configMap列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: toCells(configMapList.Items),
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

	//将[]DataCell类型的configmap列表转为v1.configmap列表
	configMaps := fromCells(data.GenericDataList)
	items := make([]*v1.ConfigMap, len(configMaps))
	for _, v := range configMaps {
		items = append(items, &v)
	}
	return &core.GetConfigMapsResp{
		Items: items,
		Total: int64(total),
	}, nil
}
