package apiv1

import (
	"keydol/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ListUserReq struct {
	g.Meta   `path:"/user/list" tags:"ListUser" method:"get" summary:"Get list of user"`
	Username string `json:"username"`
	Id       string `json:"id"`
}
type ListUserRes entity.User
