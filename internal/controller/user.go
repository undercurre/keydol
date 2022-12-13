package controller

import (
	"context"

	"keydol/apiv1"
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

func (h *cUser) Register(ctx context.Context, req *apiv1.RegisterReq) (err error) {
	return service.User().Register(ctx, req)
}
