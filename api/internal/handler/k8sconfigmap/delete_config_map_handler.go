package k8sconfigmap

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sconfigmap"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route delete /k8sconfigmap/delete_config_map k8sconfigmap DeleteConfigMap
//
// deleteConfigMap
//
// deleteConfigMap
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: DeleteConfigMapReq
//
// Responses:
//  200: DeleteConfigMapResp

func DeleteConfigMapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteConfigMapReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sconfigmap.NewDeleteConfigMapLogic(r.Context(), svcCtx)
		resp, err := l.DeleteConfigMap(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
