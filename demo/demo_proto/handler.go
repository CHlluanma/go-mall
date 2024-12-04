package main

import (
	"context"

	"github.com/chhz0/go-mall-kitex/demo/demo_proto/biz/service"
	pdapi "github.com/chhz0/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *pdapi.Request) (resp *pdapi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
