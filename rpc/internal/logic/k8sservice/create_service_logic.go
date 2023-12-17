package k8sservice

import (
	"context"
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
	service := &Service{}
	resp, err := service.CreateService(l, in)
	return resp, err
}
