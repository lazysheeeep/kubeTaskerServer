package k8sdeployment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sdeployment"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sdeployment/get_deployments k8sdeployment GetDeployments
//
// getDeployments
//
// getDeployments
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetDeploymentsReq
//
// Responses:
//  200: GetDeploymentsResp

func GetDeploymentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDeploymentsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sdeployment.NewGetDeploymentsLogic(r.Context(), svcCtx)
		resp, err := l.GetDeployments(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
