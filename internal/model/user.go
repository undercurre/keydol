package model

type UserRegisterInput struct {
	Username string
	Password string
	Email    string
	Phone    string
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
	UserId int `json:"id"`
}
