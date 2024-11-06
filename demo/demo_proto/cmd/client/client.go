package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/conf"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi/echo"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdResolver(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}

	c, err := echo.NewClient("demo_proto", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		// client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		panic(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	res, err := c.Echo(ctx, &pdapi.Request{Message: "hello"})
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("bizErr: %#v\n", bizErr)
		}
		klog.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
}
