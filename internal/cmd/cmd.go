package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/halalala222/gf-learning/internal/controller/ncu_home"
	"github.com/halalala222/gf-learning/internal/controller/user"
	"github.com/halalala222/gf-learning/internal/controller/wallet"
	"github.com/halalala222/gf-learning/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.New().Login,
				)
				group.Group("/ncuhome", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						ncu_home.New().CreatePerson,
					)
				})
				group.Group("/wallet", func(group *ghttp.RouterGroup) {
					group.Bind(
						wallet.New().Login,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
