package handler

import (
	"net/http"

	"github.com/Junhiee/lebron/apps/app/internal/logic"
	"github.com/Junhiee/lebron/apps/app/internal/svc"
	"github.com/Junhiee/lebron/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 推荐商品列表
func RecommendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRecommendLogic(r.Context(), svcCtx)
		resp, err := l.Recommend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
