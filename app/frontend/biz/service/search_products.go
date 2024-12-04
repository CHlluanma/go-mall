package service

import (
	"context"

	product "github.com/chhz0/go-mall-kitex/app/frontend/hertz_gen/frontend/product"
	"github.com/chhz0/go-mall-kitex/app/frontend/infra/rpc"
	rpcproduct "github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	products, err := rpc.ProductCatalogClient.SearchProducts(h.Context, &rpcproduct.SearchProductReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"items": products.Results,
		"q":     req.Q,
	}, nil
}
