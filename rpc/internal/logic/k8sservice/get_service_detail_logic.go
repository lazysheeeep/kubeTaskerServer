package k8sservice

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceDetailLogic {
	return &GetServiceDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceDetailLogic) GetServiceDetail(in *core.GetServiceDetailReq) (*core.GetServiceDetailResp, error) {
	// todo: add your logic here and delete this line
	service := &Service{}
	resp, _ := service.GetServiceDetail(l, in)
	return resp, nil
}
