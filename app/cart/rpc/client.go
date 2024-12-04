package rpc

import (
	"sync"

	"github.com/chhz0/go-mall-kitex/app/cart/conf"
	cartUtils "github.com/chhz0/go-mall-kitex/app/cart/utils"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	ProductCatalogClient productcatalogservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductCatalogClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
