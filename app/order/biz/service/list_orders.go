package service

import (
	"context"

	"github.com/chhz0/go-mall-kitex/app/order/biz/dal/mysql"
	"github.com/chhz0/go-mall-kitex/app/order/biz/model"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/cart"
	order "github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrdersService struct {
	ctx context.Context
} // NewListOrdersService new ListOrdersService
func NewListOrdersService(ctx context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: ctx}
}

// Run create note info
func (s *ListOrdersService) Run(req *order.ListOrdersReq) (resp *order.ListOrdersResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}

	var orders []*order.Order
	for _, l := range list {
		var items []*order.OrderItem
		for _, oi := range l.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			CreatedAt: int32(l.CreatedAt.Unix()),
			OrderId:   l.OrderId,
			UserId:    l.UserId,
			Email:     l.Consignee.Email,
			Address: &order.Address{
				StreetAddress: l.Consignee.StreetAddress,
				City:          l.Consignee.City,
				Country:       l.Consignee.Country,
				State:         l.Consignee.State,
				ZipCode:       l.Consignee.ZipCode,
			},
			Items: items,
		})
	}

	resp = &order.ListOrdersResp{
		Orders: orders,
	}
	return
}
