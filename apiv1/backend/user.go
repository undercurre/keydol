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

type UserListReq struct {
	g.Meta   `path:"/backend/user/list" tags:"User" method:"get" summary:"List of users"`
	Username string `json:"username"`
	Id       string `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	CommonPaginationReq
}

type UserListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
