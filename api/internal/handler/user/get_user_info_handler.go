package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/kubeTasker/kubeTaskerServer/api/internal/logic/user"
	"github.com/kubeTasker/kubeTaskerServer/api/internal/svc"
)

// swagger:route get /user/info user GetUserInfo
//
// Get user basic information | 获取用户基本信息
//
// Get user basic information | 获取用户基本信息
//
// Responses:
//  200: UserBaseIDInfoResp

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
