package service

import (
	"context"
	"testing"

	pdapi "github.com/CHlluanma/go-mall-kitex/demo/demo_proto/kitex_gen/pdapi"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	req := &pdapi.Request{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
