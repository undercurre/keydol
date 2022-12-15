package controller

import (
	"context"

	"keydol/apiv1"
	model "keydol/internal/model/do"
	"keydol/internal/service"
)

var (
	UserController = cUser{}
)

type cUser struct{}

func (h *cUser) List(ctx context.Context, req *apiv1.ListUserReq) (res []*apiv1.ListUserRes, err error) {
	list, err := service.User().List(ctx, req)
	return list, err
}

func (h *cUser) Register(ctx context.Context, req *apiv1.RegisterReq) (res *apiv1.RegisterRes, err error) {
	err = service.User().Register(ctx, model.UserCreateInput{
		Password: req.Password,
		Username: req.Username,
		Email:    req.Email,
	})
	return
}
