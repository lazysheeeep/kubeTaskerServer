package k8sdeployment

import (
	"context"
	"github.com/kubeTasker/kubeTaskerServer/rpc/types/core"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeployNumPerNpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeployNumPerNpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeployNumPerNpLogic {
	return &GetDeployNumPerNpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetDeployNumPerNpLogic) GetDeployNumPerNp(req *types.GetDeployNumPerNpReq) (resp *types.GetDeployNumPerNpResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.CoreRpc.GetDeployNumPerNp(l.ctx, &core.GetDeployNumPerNpReq{})
	if err != nil {
		return nil, err
	}
	deploysNps := make([]*types.DeploysNp, len(result.DeploysNps))
	for _, v := range result.DeploysNps {
		deploysNps = append(deploysNps, &types.DeploysNp{
			Namespace: v.Namespace,
			DeployNum: v.DeployNum,
		})
	}
	return &types.GetDeployNumPerNpResp{
		DeploysNps: deploysNps,
	}, err
}
