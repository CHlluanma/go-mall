package rpc

import (
	"sync"

	"github.com/CHlluanma/go-mall-kitex/app/frontend/conf"
	frontendUtils "github.com/CHlluanma/go-mall-kitex/app/frontend/utils"
	"github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().Hertz.RegistryAddress)
	frontendUtils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
