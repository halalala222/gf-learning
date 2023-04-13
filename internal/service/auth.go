package service

import (
	"context"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/halalala222/gf-learning/internal/model"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "github.com/halalala222/demo",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "username",
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   WalletLoginAuthenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserLoginInput
	)
	err := r.Parse(&in)
	if err != nil {
		return nil, err
	}
	login, err := User().Login(ctx, in)
	if err == nil {
		return login, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

func WalletLoginAuthenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in *model.WalletLoginInput
	)
	err := r.Parse(in)
	if err != nil {
		return nil, err
	}
	userId, err := Wallet().Login(ctx, in)
	if err != nil {
		return userId, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

func Unauthorized(ctx context.Context, code int, msg string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  msg,
	})
	r.ExitAll()
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	user := data.(*model.UserLoginOutput)
	if user == nil {
		return nil
	}
	return jwt.MapClaims{
		"username": user.Username,
	}
}

func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}
