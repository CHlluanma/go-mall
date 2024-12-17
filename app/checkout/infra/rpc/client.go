package rpc

import (
	"sync"

	"github.com/chhz0/go-mall-kitex/app/checkout/conf"
	"github.com/chhz0/go-mall-kitex/common/clientsuite"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/order/orderservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client

	once         sync.Once
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress

	err error
)

func Init() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr[0],
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr[0],
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr[0],
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr[0],
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
