package k8snamespace

import (
	"context"
	"errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNamespacesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNamespacesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNamespacesLogic {
	return &GetNamespacesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNamespacesLogic) GetNamespaces(in *core.GetNamespacesReq) (*core.GetNamespacesResp, error) {
	// todo: add your logic here and delete this line
	namespaceList, err := l.svcCtx.K8s.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Logger.Error(errors.New("获取Namespace列表失败, " + err.Error()))
		return nil, errors.New("获取Namespace列表失败, " + err.Error())
	}
	selectableData := &dataSelector{
		GenericDataList: toCells(namespaceList.Items),
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

	//将[]DataCell类型的namespace列表转为v1.namespace列表
	namespaces := fromCells(data.GenericDataList)
	items := make([]*corev1.Namespace, len(namespaces))
	for _, item := range namespaces {
		items = append(items, &item)
	}

	return &core.GetNamespacesResp{
		Msg: "成功！",
		Data: &core.GetNamespacesData{
			Items: items,
			Total: int64(total),
		},
	}, nil
}
