package k8sservice

import (
	"context"
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
	service := &Service{}
	resp, err := service.GetServices(l, in)
	return resp, err
}
