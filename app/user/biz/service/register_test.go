package service

import (
	"context"
	"testing"

	"github.com/CHlluanma/go-mall-kitex/app/user/biz/dal/mysql"
	user "github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test@gmail.com",
		Password:        "12345678",
		PasswordConfirm: "12345678",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
