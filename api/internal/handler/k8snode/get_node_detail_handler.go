package k8snode

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/k8snode"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /k8snode/get_node_detail k8snode GetNodeDetail
//
// getNodeDetail
//
// getNodeDetail
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: GetNodeDetailReq
//
// Responses:
//  200: GetNodeDetailResp

func GetNodeDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNodeDetailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := k8snode.NewGetNodeDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetNodeDetail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
