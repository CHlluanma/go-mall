package service

import (
	"context"

	"github.com/chhz0/go-mall-kitex/app/user/biz/dal/mysql"
	"github.com/chhz0/go-mall-kitex/app/user/biz/model"
	user "github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/user"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password or password_confirm is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("hash password failed")
	}
	newuser := &model.User{
		Email:    req.Email,
		Password: string(passwordHashed),
	}

	err = model.Create(context.Background(), mysql.DB, newuser)
	if err != nil {
		return nil, errors.New("create user failed")
	}

	return &user.RegisterResp{UserId: int32(newuser.ID)}, nil
}
