package k8sdeployment

import (
	"context"
	"encoding/json"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartDeploymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRestartDeploymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartDeploymentLogic {
	return &RestartDeploymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重启deployment
func (l *RestartDeploymentLogic) RestartDeployment(in *core.RestartDeploymentReq) (*core.RestartDeploymentResp, error) {
	// todo: add your logic here and delete this line
	//此功能等同于一下kubectl命令
	//kubectl deployment ${service} -p \
	//'{"spec":{"template":{"spec":{"containers":[{"name":"'"${service}"'","env":[{"name":"RESTART_","value":"'$(date +%s)'"}]}]}}}}'

	//使用patchData Map组装数据
	patchData := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{"name": in.DeploymentName,
							"env": []map[string]string{{
								"name":  "RESTART_",
								"value": strconv.FormatInt(time.Now().Unix(), 10),
							}},
						},
					},
				},
			},
		},
	}
	//序列化为字节，因为patch方法只接收字节类型参数
	patchByte, err := json.Marshal(patchData)
	if err != nil {
		l.Error(errors.New("json序列化失败, " + err.Error()))
		return nil, errors.New("json序列化失败, " + err.Error())
	}
	//调用patch方法更新deployment
	_, err = l.svcCtx.K8s.AppsV1().Deployments(in.Namespace).Patch(context.TODO(), in.DeploymentName, "application/strategic-merge-patch+json", patchByte, metav1.PatchOptions{})
	if err != nil {
		l.Error(errors.New("重启Deployment失败, " + err.Error()))
		return nil, errors.New("重启Deployment失败, " + err.Error())
	}

	return nil, nil
}
