package k8sPod

import (
	"bytes"
	"context"
	"errors"
	"io"
	corev1 "k8s.io/api/core/v1"

	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPodLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPodLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPodLogLogic {
	return &GetPodLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var podLogTailLine = 200

func (l *GetPodLogLogic) GetPodLog(in *core.GetPodLogReq) (*core.GetPodLogResp, error) {
	// todo: add your logic here and delete this line
	//设置日志配置,容器名,获取内容的配置
	lineLimit := int64(podLogTailLine)
	option := &corev1.PodLogOptions{
		Container: in.ContainerName,
		TailLines: &lineLimit,
	}
	// 获取一个request实例
	req := l.svcCtx.K8s.CoreV1().Pods(in.Namespace).GetLogs(in.PodName, option)
	// 发起stream连接,获取到Response.body
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		l.Error("更新Pod失败," + err.Error())
		return &core.GetPodLogResp{
			Log: "",
		}, errors.New("更新Pod失败," + err.Error())
	}
	defer func(podLogs io.ReadCloser) {
		err := podLogs.Close()
		if err != nil {
			logx.Error("Failed to close the pod log stream")
		}
	}(podLogs)
	// 将Response.body 写入到缓存区,目的为了转换成string类型
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		l.Error("复制podLog失败," + err.Error())
		return &core.GetPodLogResp{
			Log: "",
		}, errors.New("复制podLog失败," + err.Error())
	}
	return &core.GetPodLogResp{
		Log: buf.String(),
	}, nil
}
