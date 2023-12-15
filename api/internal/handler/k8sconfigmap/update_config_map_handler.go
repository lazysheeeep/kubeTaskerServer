package k8sconfigmap

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sconfigmap"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route put /k8sconfigmap/update_config_map k8sconfigmap UpdateConfigMap
//
// updateConfigMap
//
// updateConfigMap
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpdateConfigMapReq
//
// Responses:
//  200: UpdateConfigMapResp

func UpdateConfigMapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateConfigMapReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sconfigmap.NewUpdateConfigMapLogic(r.Context(), svcCtx)
		resp, err := l.UpdateConfigMap(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
