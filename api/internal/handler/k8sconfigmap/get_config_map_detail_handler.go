package k8sconfigmap

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sconfigmap"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sconfigmap/get_config_map_detail k8sconfigmap GetConfigMapDetail
//
// getConfigMapDetail
//
// getConfigMapDetail
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetConfigMapDetailReq
//
// Responses:
//  200: GetConfigMapDetailResp

func GetConfigMapDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetConfigMapDetailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sconfigmap.NewGetConfigMapDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetConfigMapDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
