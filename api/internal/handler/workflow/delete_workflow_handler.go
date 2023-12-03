package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route delete /workflow/delete_workflow/{namespace}/{name} workflow DeleteWorkflow
//
// DeleteWorkflow
//
// DeleteWorkflow
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowDeleteRequest
//
// Responses:
//  200: WorkflowDeleteResponse

func DeleteWorkflowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowDeleteRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewDeleteWorkflowLogic(r.Context(), svcCtx)
		resp, err := l.DeleteWorkflow(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
