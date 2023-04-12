package cmd

import (
	"context"
	"github.com/halalala222/gf-learning/internal/controller/ncu_home"
	"github.com/halalala222/gf-learning/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().Auth)
				group.Bind(
					ncu_home.New(),
				)
			})
			s.Run()
			return nil
		},
	}
)
