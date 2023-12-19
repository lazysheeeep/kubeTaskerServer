package k8snamespace

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8snamespace"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8s_namespace/get_namespaces k8snamespace GetNamespaces
//
// getNamespaces
//
// getNamespaces
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetNamespacesReq
//
// Responses:
//  200: GetNamespacesResp

func GetNamespacesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNamespacesReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8snamespace.NewGetNamespacesLogic(r.Context(), svcCtx)
		resp, err := l.GetNamespaces(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
