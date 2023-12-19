package k8spod

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8spod"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8spod/get_pod_detail k8spod GetPodDetail
//
// getPodDetail
//
// getPodDetail
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetPodDetailReq
//
// Responses:
//  200: GetPodDetailResp

func GetPodDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPodDetailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8spod.NewGetPodDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetPodDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
