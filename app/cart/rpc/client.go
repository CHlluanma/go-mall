package rpc

import (
	"sync"

	"github.com/chhz0/go-mall-kitex/app/cart/conf"
	cartUtils "github.com/chhz0/go-mall-kitex/app/cart/utils"
	"github.com/chhz0/go-mall-kitex/common/clientsuite"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductCatalogClient productcatalogservice.Client
	once sync.Once

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress
	err          error
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr[0],
		}),
	}

	ProductCatalogClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
