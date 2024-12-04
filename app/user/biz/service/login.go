package service

import (
	"context"
	"errors"

	"github.com/chhz0/go-mall-kitex/app/user/biz/dal/mysql"
	"github.com/chhz0/go-mall-kitex/app/user/biz/model"
	user "github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	row, err := model.GetByEmail(context.Background(), mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password is wrong")
	}
	resp = &user.LoginResp{UserId: int32(row.ID)}
	return resp, nil
}
