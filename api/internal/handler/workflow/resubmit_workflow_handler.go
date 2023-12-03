package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route put /workflow/resubmit_workflow/{namespace}/{name}/resubmit workflow ResubmitWorkflow
//
// ResubmitWorkflow
//
// ResubmitWorkflow
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowResubmitRequest
//
// Responses:
//  200: WorkflowRespond

func ResubmitWorkflowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowResubmitRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewResubmitWorkflowLogic(r.Context(), svcCtx)
		resp, err := l.ResubmitWorkflow(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
