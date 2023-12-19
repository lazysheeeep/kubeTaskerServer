package k8sdeployment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sdeployment"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sdeployment/scale_deployment k8sdeployment ScaleDeployment
//
// scaleDeployment
//
// scaleDeployment
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ScaleDeploymentReq
//
// Responses:
//  200: ScaleDeploymentResp

func ScaleDeploymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ScaleDeploymentReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sdeployment.NewScaleDeploymentLogic(r.Context(), svcCtx)
		resp, err := l.ScaleDeployment(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
