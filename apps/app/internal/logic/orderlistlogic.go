package logic

import (
	"context"
	"strconv"
	"strings"

	"github.com/Junhiee/lebron/apps/app/internal/svc"
	"github.com/Junhiee/lebron/apps/app/internal/types"
	"github.com/Junhiee/lebron/apps/order/rpc/order"
	"github.com/Junhiee/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单列表
func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderListResponse, err error) {
	// todo: add your logic here and delete this line
	orderRes, err := l.svcCtx.OrderRPC.Orders(l.ctx, &order.OrdersRequest{UserId: req.UID})
	if err != nil {
		return nil, err
	}
	var pids []string
	for _, o := range orderRes.Orders {
		pids = append(pids, strconv.Itoa(int(o.Proid)))
	}

	productRes, err := l.svcCtx.ProductRPC.Products(l.ctx, &product.ProductRequest{ProductIds: strings.Join(pids, ",")})
	if err != nil {
		return nil, err
	}
	var orders []*types.Order
	for _, o := range orderRes.Orders {
		if p, ok := productRes.Products[o.Proid]; ok {
			orders = append(orders, &types.Order{
				OrderID:     o.Orderid,
				ProductName: p.Name,
			})
		}
	}
	return &types.OrderListResponse{Orders: orders}, nil

}
