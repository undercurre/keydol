package cmd

import (
	"context"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	TestServerName string
	gfToken        *gtoken.GfToken
	Main           = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server(TestServerName)

			// 不认证接口
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// 调试路由
				group.ALL("/backend/user/register", func(r *ghttp.Request) {
					r.Response.WriteJson(gtoken.Succ("hello"))
				})
			})
			MultiLogin, err := g.Cfg().Get(ctx, "gToken.MultiLogin")
			// 认证接口
			loginFunc := Login
			// 启动gtoken
			gfToken = &gtoken.GfToken{
				ServerName:       TestServerName,
				LoginPath:        "/login",
				LoginBeforeFunc:  loginFunc,
				LogoutPath:       "/user/logout",
				AuthExcludePaths: g.SliceStr{"/user/list", "/backend/user/list"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				MultiLogin:       MultiLogin.Bool(),
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.ALL("/backend/user/list", func(r *ghttp.Request) {
					r.Response.WriteJson(gtoken.Succ("system user info"))
				})
			})
			s.Run()
			return nil
		},
	}
)

func Login(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	password := r.Get("password").String()

	if username == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return username, "mieye"
}
