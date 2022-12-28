package cmd

import (
	"context"
	"keybol/internal/consts"
	"keybol/internal/controller"
	"keybol/internal/dao"
	"keybol/internal/middleware"
	"keybol/internal/model/entity"
	"keybol/utility"
	"strconv"

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
				group.Middleware(middleware.MiddlewareCORS)
				// 调试路由
			})
			// 认证接口
			loginFunc := Login
			// 启动gtoken
			gfToken = &gtoken.GfToken{
				CacheMode:        1,
				ServerName:       TestServerName,
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  loginFunc,
				LogoutPath:       "/backend/user/logout",
				AuthPaths:        g.SliceStr{"/backend/user/list"},
				AuthExcludePaths: g.SliceStr{"/backend/user/register"},
				MultiLogin:       true,
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.MiddlewareCORS)
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					controller.User,
				)
			})
			s.Run()
			return nil
		},
	}
)

func Login(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if username == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	User := entity.User{}
	err := dao.User.Ctx(ctx).Where("username", username).Scan(&User)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, User.Usersalt) != User.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(User.Id), User
}
