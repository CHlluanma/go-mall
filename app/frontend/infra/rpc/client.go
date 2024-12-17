package rpc

import (
	"sync"

	"github.com/chhz0/go-mall-kitex/app/frontend/conf"
	frontendUtils "github.com/chhz0/go-mall-kitex/app/frontend/utils"
	"github.com/chhz0/go-mall-kitex/common/clientsuite"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/order/orderservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient           userservice.Client
	ProductCatalogClient productcatalogservice.Client
	CartClient           cartservice.Client
	CheckoutClient       checkoutservice.Client
	OrderClient          orderservice.Client

	once         sync.Once
	serviceName  = frontendUtils.ServiceName
	registryAddr = conf.GetConf().Hertz.RegistryAddress
	err          error
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr[0],
		}),
	}

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr[0],
		}),
	}

	ProductCatalogClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr[0],
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr[0],
		}),
	}

	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr[0],
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}
