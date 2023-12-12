package k8sPod

import (
	"context"
	"errors"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	v1 "k8s.io/api/core/v1"

	"github.com/zeromicro/go-zero/core/logx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GetPodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodsLogic {
	return &GetPodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// K8sPod management
func (l *GetPodsLogic) GetPods(in *core.GetPodsReq) (*core.GetPodsResp, error) {
	// todo: add your logic here and delete this line
	podList, err := l.svcCtx.K8s.CoreV1().Pods(in.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Error("获取Pod列表失败," + err.Error())
		// 返回给上一层,最终返回给前端,前端捕获到后打印出来
		return nil, errors.New("获取Pod列表失败," + err.Error())
	}
	// 实例化dataselector结构体,组装数据
	selectableData := &dataSelector{
		GenericDataList: toCells(podList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: in.FilterName},
			Paginate: &PaginateQuery{
				Limit: int(in.Limit),
				Page:  int(in.Page),
			},
		},
	}
	// 先过滤
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	// 排序和分页
	data := filtered.Sort().Paginate()
	println("Pod total: ", total)
	// 将DataCell类型转成Pod
	pods := fromCells(data.GenericDataList)
	items := make([]*v1.Pod, len(pods))
	for _, value := range pods {
		items = append(items, &value)
	}
	return &core.GetPodsResp{
		Items: items,
		Total: int32(total),
	}, nil
}
