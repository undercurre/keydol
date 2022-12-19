package user

import (
	"context"
	"keybol/internal/dao"
	"keybol/internal/model"
	"keybol/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.UserRegisterInput) (out model.UserRegisterOutput, err error) {
	// 不允许HTML代码
	// if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	// 	return out, err
	// }
	//处理加密盐和密码的逻辑
	// UserSalt := grand.S(10)
	// in.Password = utility.EncryptPassword(in.Password, UserSalt)
	// in.UserSalt = UserSalt
	//插入数据返回id
	lastInsertID, err := dao.User.Ctx(ctx).Data(g.Map{
		"username":   in.Username,
		"password":   in.Password,
		"email":      in.Email,
		"phone":      in.Phone,
		"createTime": gdb.Raw("now()"),
		"updateTime": gdb.Raw("now()"),
	}).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserRegisterOutput{UserId: int(lastInsertID)}, err
}
