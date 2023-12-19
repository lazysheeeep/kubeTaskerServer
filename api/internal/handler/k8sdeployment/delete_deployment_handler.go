package k8sdeployment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sdeployment"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sdeployment/delete_deployment k8sdeployment DeleteDeployment
//
// deleteDeployment
//
// deleteDeployment
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DeleteDeploymentReq
//
// Responses:
//  200: DeleteDeploymentResp

func DeleteDeploymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteDeploymentReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sdeployment.NewDeleteDeploymentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDeployment(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
