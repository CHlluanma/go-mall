package main

import (
	"context"
	"fmt"

	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/conf"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi"
	"github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}

	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		panic(err)
	}

	res, err := c.Echo(context.Background(), &pdapi.Request{Message: "hello"})
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Printf("res: %v", res)
}
