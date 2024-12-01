package service

import (
	"context"
	"time"

	common "github.com/CHlluanma/go-mall-kitex/app/frontend/hertz_gen/frontend/common"
	"github.com/CHlluanma/go-mall-kitex/app/frontend/infra/rpc"
	"github.com/CHlluanma/go-mall-kitex/app/frontend/types"
	frontendUtils "github.com/CHlluanma/go-mall-kitex/app/frontend/utils"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/order"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrders(h.Context, &order.ListOrdersReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {
		var total float32
		var items []types.OrderItem

		for _, item := range v.Items {
			total += item.Cost
			i := item.Item
			productResp, err := rpc.ProductCatalogClient.GetProduct(h.Context, &product.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}
			p := productResp.Product

			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Cost:        item.Cost,
				Qty:         i.Quantity,
			})
		}
		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
