package service

import (
	"context"

	auth "github.com/chhz0/go-mall-kitex/app/frontend/hertz_gen/frontend/auth"
	"github.com/chhz0/go-mall-kitex/app/frontend/infra/rpc"
	"github.com/chhz0/go-mall-kitex/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	rpcResp, err := rpc.UserClient.Login(h.Context,
		&user.LoginReq{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		return "", err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", rpcResp.UserId)
	err = session.Save()
	if err != nil {
		return "", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return
}
