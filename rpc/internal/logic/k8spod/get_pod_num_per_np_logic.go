package k8spod

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodNumPerNpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodNumPerNpLogic {
	return &GetPodNumPerNpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPodNumPerNpLogic) GetPodNumPerNp(in *core.GetPodNumPerNpReq) (*core.GetPodNumPerNpResp, error) {
	// todo: add your logic here and delete this line
	namespaceList, err := l.svcCtx.K8s.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	podsNps := make([]*core.PodsNp, 0)
	for _, namespace := range namespaceList.Items {
		//获取pod列表
		podList, err := l.svcCtx.K8s.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		//组装数据
		podsNp := &core.PodsNp{
			Namespace: namespace.Name,
			PodNum:    int64(len(podList.Items)),
		}
		//添加到podsNps数组中
		podsNps = append(podsNps, podsNp)
	}
	return &core.GetPodNumPerNpResp{
		PodsNps: podsNps,
	}, nil
}
