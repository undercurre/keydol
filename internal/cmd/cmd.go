package cmd

import (
	"context"
	"fmt"
	"keybol/internal/consts"
	"keybol/internal/controller"
	"keybol/internal/dao"
	"keybol/internal/middleware"
	"keybol/internal/model/entity"
	"keybol/utility"
	"strconv"

	"net/url"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
)

type Response struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

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
			loginBackFunc := LoginBack
			// 启动gtoken
			gfToken = &gtoken.GfToken{
				CacheMode:        1,
				ServerName:       TestServerName,
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  loginFunc,
				LoginAfterFunc:   loginBackFunc,
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
					controller.Mission,
				)
			})
			s.Run()
			return nil
		},
	}
)

func Login(r *ghttp.Request) (string, interface{}) {
	loginType := r.Get("type").String()
	loginCode := r.Get("code").String()
	username := r.Get("username").String()
	password := r.Get("password").String()

	ctx := context.TODO()

	User := entity.User{}

	// 账密
	if loginType == "primary" || loginType == "" {
		if username == "" || password == "" {
			r.Response.WriteJson(gtoken.Fail("账号或密码为空."))
			r.ExitAll()
		}

		// 查询不到
		err := dao.User.Ctx(ctx).Where("username", username).Scan(&User)
		if err != nil {
			r.Response.WriteJson(gtoken.Fail("没有这个账号."))
			r.ExitAll()
		}

		// 密码错误
		if utility.EncryptPassword(password, User.Usersalt) != User.Password {
			r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
			r.ExitAll()
		}
	}

	// 微信
	if loginType == "wechat" {
		if loginCode == "" {
			r.Response.WriteJson(gtoken.Fail("账号或密码为空."))
		}

		client := g.Client()

		var wechatBaseURL = "https://api.weixin.qq.com/sns/jscode2session"

		// 构建一个 URL 对象
		u, err := url.Parse(wechatBaseURL)
		if err != nil {
			panic(err)
		}

		// 设置查询参数
		q := u.Query()
		q.Set("appid", "wx7283eac281febaaf")
		q.Set("secret", "d67a5e24c9af975b5cde81aa64d64613")
		q.Set("js_code", "0a1XhIkl2rlI3b4YA6nl2kpTKr1XhIk-")
		q.Set("grant_type", "authorization_code")
		u.RawQuery = q.Encode()
		urlStr := u.String()

		resp, err := client.Get(ctx, urlStr)

		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("status:", resp.Status)
		fmt.Println("body:", resp.Body)

		// 查询
		// err := dao.
	}

	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(User.Id), User
}

func LoginBack(r *ghttp.Request, respData gtoken.Resp) {
	//获得登录用户id
	userKey := respData.GetString("userKey")
	userId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
	//根据id获得登录用户其他信息
	userInfo := entity.User{}
	err := dao.User.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return
	}
	r.Response.WriteJson(gtoken.Succ(g.Map{
		"token": respData.GetString("token"),
		"info": g.Map{
			"userId":   userInfo.Id,
			"username": userInfo.Username,
			"email":    userInfo.Email,
			"phone":    userInfo.Phone,
			"roleIds":  userInfo.RoleIds,
		},
	}))
}
