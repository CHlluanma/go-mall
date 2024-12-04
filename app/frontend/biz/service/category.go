package service

import (
	"context"

	category "github.com/chhz0/go-mall-kitex/app/frontend/hertz_gen/frontend/category"
	"github.com/chhz0/go-mall-kitex/app/frontend/infra/rpc"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductCatalogClient.ListProducts(h.Context, &product.ListProductsReq{
		CategoryName: req.Category,
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
