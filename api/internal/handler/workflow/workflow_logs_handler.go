package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route get /workflow/workflow_logs/{namespace}/{name}/log workflow WorkflowLogs
//
// WorkflowLogs
//
// WorkflowLogs
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowLogRequest
//
// Responses:
//  200: LogEntry

func WorkflowLogsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowLogRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewWorkflowLogsLogic(r.Context(), svcCtx)
		resp, err := l.WorkflowLogs(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}