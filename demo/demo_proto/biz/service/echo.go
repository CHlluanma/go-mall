package service

import (
	"context"

	pdapi "github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi"
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

	return &pdapi.Response{Message: req.Message}, nil
}
