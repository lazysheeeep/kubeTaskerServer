package k8snamespace

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8snamespace"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8s_namespace/get_namespace_detail k8snamespace GetNamespaceDetail
//
// getNamespaceDetail
//
// getNamespaceDetail
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetNamespaceDetailReq
//
// Responses:
//  200: GetNamespaceDetailResp

func GetNamespaceDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNamespaceDetailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8snamespace.NewGetNamespaceDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetNamespaceDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
