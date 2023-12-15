package k8sservice

import (
	"context"
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServicesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServicesLogic {
	return &GetServicesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取service列表，支持过滤、排序、分页
func (l *GetServicesLogic) GetServices(in *core.GetServicesReq) (*core.GetServicesResp, error) {
	// todo: add your logic here and delete this line
	//获取serviceList类型的service列表
	serviceList, err := l.svcCtx.K8s.CoreV1().Services(in.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Error(errors.New("获取Service列表失败, " + err.Error()))
		return nil, errors.New("获取Service列表失败, " + err.Error())
	}
	//将serviceList中的service列表(Items)，放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: toCells(serviceList.Items),
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

	//将[]DataCell类型的service列表转为v1.service列表
	services := fromCells(data.GenericDataList)
	items := make([]*v1.Service, 0)
	for _, v := range services {
		items = append(items, &v)
	}
	return &core.GetServicesResp{
		Items: items,
		Total: int64(total),
	}, nil
}
