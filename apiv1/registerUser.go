package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 用户列表

type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"get" tags:"UserService" method:"post" summary:"Register a new user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterRes struct{}
