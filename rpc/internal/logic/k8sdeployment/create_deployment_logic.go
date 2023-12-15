package k8sdeployment

import (
	"context"
	"errors"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeploymentLogic {
	return &CreateDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建deployment,接收DeployCreate对象
func (l *CreateDeploymentLogic) CreateDeployment(in *core.CreateDeploymentReq) (*core.CreateDeploymentResp, error) {
	//将data中的属性组装成appsv1.Deployment对象,并将入参数据放入
	deployment := &appsv1.Deployment{
		//ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      in.Name,
			Namespace: in.Namespace,
			Labels:    in.Label,
		},
		//Spec中定义副本数、选择器、以及pod属性
		Spec: appsv1.DeploymentSpec{
			Replicas: &in.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: in.Label,
			},
			Template: corev1.PodTemplateSpec{
				//定义pod名和标签
				ObjectMeta: metav1.ObjectMeta{
					Name:   in.Name,
					Labels: in.Label,
				},
				//定义容器名、镜像和端口
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  in.Name,
							Image: in.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: in.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
		//Status定义资源的运行状态，这里由于是新建，传入空的appsv1.DeploymentStatus{}对象即可
		Status: appsv1.DeploymentStatus{},
	}
	//判断是否打开健康检查功能，若打开，则定义ReadinessProbe和LivenessProbe
	if in.HealthCheck {
		//设置第一个容器的ReadinessProbe，因为我们pod中只有一个容器，所以直接使用index 0即可
		//若pod中有多个容器，则这里需要使用for循环去定义了
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: in.HealthPath,
					//intstr.IntOrString的作用是端口可以定义为整型，也可以定义为字符串
					//Type=0则表示表示该结构体实例内的数据为整型，转json时只使用IntVal的数据
					//Type=1则表示表示该结构体实例内的数据为字符串，转json时只使用StrVal的数据
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: in.ContainerPort,
					},
				},
			},
			//初始化等待时间
			InitialDelaySeconds: 5,
			//超时时间
			TimeoutSeconds: 5,
			//执行间隔
			PeriodSeconds: 5,
		}
		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: in.HealthPath,
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: in.ContainerPort,
					},
				},
			},
			InitialDelaySeconds: 15,
			TimeoutSeconds:      5,
			PeriodSeconds:       5,
		}
		//定义容器的limit和request资源
		deployment.Spec.Template.Spec.Containers[0].Resources.Limits = map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    resource.MustParse(in.Cpu),
			corev1.ResourceMemory: resource.MustParse(in.Memory),
		}
		deployment.Spec.Template.Spec.Containers[0].Resources.Requests = map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    resource.MustParse(in.Cpu),
			corev1.ResourceMemory: resource.MustParse(in.Memory),
		}
	}
	//调用sdk创建deployment
	_, err := l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).
		Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		l.Error(errors.New("创建Deployment失败, " + err.Error()))
		return nil, errors.New("创建Deployment失败, " + err.Error())
	}
	return nil, nil
}
