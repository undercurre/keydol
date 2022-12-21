package model

import "github.com/gogf/gf/v2/os/gtime"

type UserRegisterInput struct {
	Username string
	Password string
	Email    string
	Phone    string
	RoleIds  string
}

type UserRegisterOutput struct {
	UserId int `json:"id"`
}

type UserListInput struct {
	Username string
	Email    string
	Phone    string
	Page     int `json:"page" description:"分页码"`
	Size     int `json:"size" description:"分页数量"`
}

type UserListOutput struct {
	List  []UserListOutputItem `json:"list" description:"列表"`
	Page  int                  `json:"page" description:"分页码"`
	Size  int                  `json:"size" description:"分页数量"`
	Total int                  `json:"total" description:"数据总数"`
}

type UserListOutputItem struct {
	Id         int         `json:"id"         ` //
	Username   string      `json:"username"   ` //
	Email      string      `json:"email"      ` //
	Phone      string      `json:"phone"      ` //
	Createtime *gtime.Time `json:"createtime" ` //
	Updatetime *gtime.Time `json:"updatetime" ` //
}
