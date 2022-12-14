package service

import (
	"context"
	"fmt"
	v1 "keydol/apiv1"
	model "keydol/internal/model/do"
	"keydol/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

type UserServ struct{}

var (
	insSysUser = UserServ{}
)

func User() *UserServ {
	return &insSysUser
}

func (*UserServ) List(ctx context.Context, param *v1.ListUserReq) ([]*v1.ListUserRes, error) {
	m := dao.User.Ctx(ctx).Safe(false).FieldsEx(dao.User.Columns().Password, dao.User.Columns())

	// 添加查询条件
	if param.Username != "" {
		m.WhereLike(dao.User.Columns().Username, Like(param.Username))
	}
	if param.Id != "" {
		m.WhereLike(dao.User.Columns().Id, Like(param.Id))
	}

	// 查询
	var list []*v1.ListUserRes
	err := m.Scan(&list)
	fmt.Println(list)
	return list, err
}

func (*UserServ) Register(ctx context.Context, in model.UserCreateInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(g.Map{"id": in.Id, "username": in.Username, "password": in.Password}).Insert()
	return err
}

func Like(s string) string {
	return "%" + s + "%"
}
