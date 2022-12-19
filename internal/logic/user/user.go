package user

import (
	"context"
	"keybol/internal/dao"
	"keybol/internal/model"
	"keybol/internal/model/entity"
	"keybol/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

func (s *sUser) List(ctx context.Context, in model.UserListInput) (out *model.UserListOutput, err error) {
	var (
		m = dao.User.Ctx(ctx)
	)
	out = &model.UserListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	// listModel = listModel.OrderDesc(dao.User.Columns().Username)

	// 执行查询
	var list []*entity.User
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	var (
		temp, tempErr = m.Count()
	)
	out.Total = gconv.Int(temp)
	err = tempErr
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
