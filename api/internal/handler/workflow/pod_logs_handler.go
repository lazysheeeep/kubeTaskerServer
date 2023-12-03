package workflow

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/workflow"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/types"
)

// swagger:route get /workflow/pod_logs/{namespace}/{name}/{podName}/log workflow PodLogs
//

//

//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: WorkflowLogRequest
//
// Responses:
//  200: LogEntry

func PodLogsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WorkflowLogRequest
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := workflow.NewPodLogsLogic(r.Context(), svcCtx)
		resp, err := l.PodLogs(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
