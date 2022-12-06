package apiv1

import (
	"keydol/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 用户列表

type RegisterReq struct {
	g.Meta   `path:"/user/" tags:"ListUser" method:"get" summary:"Get list of user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Id       string `json:"id"`
}

type RegisterRes entity.User
