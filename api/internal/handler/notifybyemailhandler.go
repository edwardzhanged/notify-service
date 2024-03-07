package handler

import (
	"net/http"

	"notify-service/api/internal/logic"
	"notify-service/api/internal/svc"
	"notify-service/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func NotifyByEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyByEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNotifyByEmailLogic(r.Context(), svcCtx)
		resp, err := l.NotifyByEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
