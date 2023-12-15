package k8sdeployment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sdeployment"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sdeployment/update_deployment k8sdeployment UpdateDeployment
//
// updateDeployment
//
// updateDeployment
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateDeploymentReq
//
// Responses:
//  200: UpdateDeploymentResp

func UpdateDeploymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDeploymentReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sdeployment.NewUpdateDeploymentLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDeployment(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
