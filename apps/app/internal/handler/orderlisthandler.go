package handler

import (
	"net/http"

	"github.com/Junhiee/lebron/apps/app/internal/logic"
	"github.com/Junhiee/lebron/apps/app/internal/svc"
	"github.com/Junhiee/lebron/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 订单列表
func OrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderListLogic(r.Context(), svcCtx)
		resp, err := l.OrderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
