package controller

import (
	"context"
	"keybol/apiv1/backend"
	"keybol/internal/model"
	"keybol/internal/service"
)

var User = cUser{}

type cUser struct{}

func (a *cUser) Register(ctx context.Context, req *backend.UserRegisterReq) (res *backend.UserRegisterRes, err error) {
	out, err := service.User().Register(ctx, model.UserRegisterInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &backend.UserRegisterRes{UserId: out.UserId}, nil
}

func (a *cUser) List(ctx context.Context, req *backend.UserListReq) (res *backend.UserListRes, err error) {
	list, err := service.User().List(ctx, model.UserListInput{
		Page:     req.Page,
		Size:     req.Size,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &backend.UserListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}, nil
}
