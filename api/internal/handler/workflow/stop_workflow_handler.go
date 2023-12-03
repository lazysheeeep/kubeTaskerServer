package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route put /workflow/stop_workflow/{namespace}/{name}/stop workflow StopWorkflow
//
// StopWorkflow
//
// StopWorkflow
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowStopRequest
//
// Responses:
//  200: WorkflowRespond

func StopWorkflowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowStopRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewStopWorkflowLogic(r.Context(), svcCtx)
		resp, err := l.StopWorkflow(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
