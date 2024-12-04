package rpc

import (
	"sync"

	"github.com/chhz0/go-mall-kitex/app/checkout/conf"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	// err           error
)

func Init() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
	})
}

func initCartClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initProductClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}
