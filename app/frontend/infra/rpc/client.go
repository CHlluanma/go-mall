package rpc

import (
	"sync"

	"github.com/CHlluanma/go-mall-kitex/app/frontend/conf"
	frontendUtils "github.com/CHlluanma/go-mall-kitex/app/frontend/utils"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	UserClient           userservice.Client
	ProductCatalogClient productcatalogservice.Client
	CartClient           cartservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductCatalogClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}