package service

import (
	"context"

	product "github.com/CHlluanma/go-mall-kitex/app/frontend/hertz_gen/frontend/product"
	"github.com/CHlluanma/go-mall-kitex/app/frontend/infra/rpc"
	rpcproduct "github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductCatalogClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
