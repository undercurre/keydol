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
