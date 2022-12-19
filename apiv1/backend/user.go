package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserRegisterReq struct {
	g.Meta   `path:"/backend/user/register" tags:"User" method:"post" summary:"Register a new user"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	Email    string `json:"email" v:"required#邮箱地址不能为空" dc:"邮箱地址"`
	Phone    string `json:"phone"    v:"required#手机号码不能为空" dc:"手机号码"`
}

type UserRegisterRes struct {
	UserId int `json:"id"`
}
