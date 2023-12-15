package k8sservice

import (
	"context"
	"errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建service,,接收ServiceCreate对象
func (l *CreateServiceLogic) CreateService(in *core.CreateServiceReq) (*core.CreateServiceResp, error) {
	// todo: add your logic here and delete this line
	//将data中的数据组装成corev1.Service对象
	service := &corev1.Service{
		//ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      in.Name,
			Namespace: in.Namespace,
			Labels:    in.Label,
		},
		//Spec中定义类型，端口，选择器
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(in.Type),
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Port:     int32(in.Port),
					Protocol: "TCP",
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: int32(in.ContainerPort),
					},
				},
			},
			Selector: in.Label,
		},
	}
	//默认ClusterIP,这里是判断NodePort,添加配置
	if in.NodePort != 0 && in.Type == "NodePort" {
		service.Spec.Ports[0].NodePort = int32(in.NodePort)
	}
	//创建Service
	_, err := l.svcCtx.K8s.CoreV1().Services(in.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		l.Error(errors.New("创建Service失败, " + err.Error()))
		return nil, errors.New("创建Service失败, " + err.Error())
	}

	return nil, nil
}
