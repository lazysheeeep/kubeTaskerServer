package k8sdeployment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sdeployment"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sdeployment/get_deploy_num_per_np k8sdeployment GetDeployNumPerNp
//
// getDeployNumPerNp
//
// getDeployNumPerNp
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetDeployNumPerNpReq
//
// Responses:
//  200: GetDeployNumPerNpResp

func GetDeployNumPerNpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDeployNumPerNpReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sdeployment.NewGetDeployNumPerNpLogic(r.Context(), svcCtx)
		resp, err := l.GetDeployNumPerNp(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
