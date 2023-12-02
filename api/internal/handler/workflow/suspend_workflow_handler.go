package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route put /workflow/suspend_workflow/{namespace}/{name}/suspend workflow SuspendWorkflow
//
// SuspendWorkflow
//
// SuspendWorkflow
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowSuspendRequest
//
// Responses:
//  200: WorkflowRespond

func SuspendWorkflowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowSuspendRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewSuspendWorkflowLogic(r.Context(), svcCtx)
		resp, err := l.SuspendWorkflow(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
