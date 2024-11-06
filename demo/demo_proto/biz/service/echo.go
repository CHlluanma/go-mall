package service

import (
	"context"
	"fmt"

	pdapi "github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pdapi.Request) (resp *pdapi.Response, err error) {
	// Finish your business logic.
	clientName, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	fmt.Println(clientName, ok)
	if req.Message == "error" {
		return nil, kerrors.NewGRPCBizStatusError(1004001, "client param error")
	}

	return &pdapi.Response{Message: req.Message}, nil
}
