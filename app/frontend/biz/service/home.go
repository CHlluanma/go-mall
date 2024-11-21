package service

import (
	"context"

	home "github.com/CHlluanma/go-mall-kitex/app/frontend/hertz_gen/frontend/common"
	"github.com/CHlluanma/go-mall-kitex/app/frontend/infra/rpc"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	products, err := rpc.ProductCatalogClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot Sales",
		"items": products.Products,
	}, nil
}
