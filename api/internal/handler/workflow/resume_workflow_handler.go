package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route put /workflow/resume_workflow/{namespace}/{name}/resume workflow ResumeWorkflow
//
// ResumeWorkflow
//
// ResumeWorkflow
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowResumeRequest
//
// Responses:
//  200: WorkflowRespond

func ResumeWorkflowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowResumeRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewResumeWorkflowLogic(r.Context(), svcCtx)
		resp, err := l.ResumeWorkflow(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
