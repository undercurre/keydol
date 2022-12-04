package router

import (
	"context"
	"keydol/internal/controller"
	"keydol/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func SetupRouter(ctx context.Context, s *ghttp.Server) {
	/* 用户管理 */
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.HandleResponse)
		group.Bind(controller.UserController)
	})
}
