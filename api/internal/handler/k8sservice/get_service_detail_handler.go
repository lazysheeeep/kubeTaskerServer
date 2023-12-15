package k8sservice

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8sservice"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8sservice/get_service_detail k8sservice GetServiceDetail
//
// getServiceDetail
//
// getServiceDetail
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetServiceDetailReq
//
// Responses:
//  200: GetServiceDetailResp

func GetServiceDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetServiceDetailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8sservice.NewGetServiceDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetServiceDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
