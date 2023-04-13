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
					wallet.New().Login,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/ncuhome": ncu_home.New().CreatePerson,
					})
				})

			})
			s.Run()
			return nil
		},
	}
)
