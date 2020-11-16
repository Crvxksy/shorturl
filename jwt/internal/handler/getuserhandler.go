package handler

import (
	"net/http"

	"shorturl/jwt/internal/logic"
	"shorturl/jwt/internal/svc"
	"shorturl/jwt/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), ctx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
