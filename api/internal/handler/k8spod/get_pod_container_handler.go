package k8spod

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8spod"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8spod/get_pod_container k8spod GetPodContainer
//
// getPodContainer
//
// getPodContainer
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetPodContainerReq
//
// Responses:
//  200: GetPodContainerResp

func GetPodContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPodContainerReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8spod.NewGetPodContainerLogic(r.Context(), svcCtx)
		resp, err := l.GetPodContainer(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
