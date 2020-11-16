package handler

import (
	"net/http"

	"shorturl/api/internal/logic"
	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

// func jwtHandler(ctx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		l := logic.NewJwtLogic(r.Context(), ctx)
// 		err := l.Jwt()
// 		if err != nil {
// 			httpx.Error(w, err)
// 		} else {
// 			httpx.Ok(w)
// 		}
// 	}
// }
func jwtHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JwtTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewJwtLogic(r.Context(), ctx)
		resp, err := l.Jwt(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
