package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route get /workflow/watch_workflows/{namespace} workflow WatchWorkflows
//
// WatchWorkflows
//
// WatchWorkflows
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WatchWorkflowsRequest
//
// Responses:
//  200: WorkflowWatchEvent

func WatchWorkflowsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WatchWorkflowsRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewWatchWorkflowsLogic(r.Context(), svcCtx)
		resp, err := l.WatchWorkflows(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
