package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route post /workflow/submit_workflow/{namespace}/submit workflow SubmitWorkflow
//
// SubmitWorkflow
//
// SubmitWorkflow
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowSubmitRequest
//
// Responses:
//  200: WorkflowRespond

func SubmitWorkflowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowSubmitRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewSubmitWorkflowLogic(r.Context(), svcCtx)
		resp, err := l.SubmitWorkflow(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
