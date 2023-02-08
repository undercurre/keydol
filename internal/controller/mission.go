package controller

import (
	"context"
	"keybol/apiv1/backend"
	"keybol/internal/model"
	"keybol/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

var Mission = cMission{}

type cMission struct{}

func (a *cMission) Create(ctx context.Context, req *backend.MissionCreateReq) (res *backend.MissionCreateRes, err error) {
	out, err := service.Mission().Create(ctx, model.MissionCreateInput{
		Missionname: req.MissionName,
		Detail:      req.Detail,
		Deadline:    gtime.New(req.Deadline),
		UserId:      req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.MissionCreateRes{MissionId: out.MissionId}, nil
}

func (a *cMission) Delete(ctx context.Context, req *backend.MissionDeleteReq) (res *backend.MissionDeleteRes, err error) {
	out, err := service.Mission().Delete(ctx, model.MissionDeleteInput{
		MissionId: req.MissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.MissionDeleteRes{MissionId: out.MissionId}, nil
}

func (a *cMission) Update(ctx context.Context, req *backend.MissionUpdateReq) (res *backend.MissionUpdateRes, err error) {
	out, err := service.Mission().Update(ctx, model.MissionUpdateInput{
		MissionName: req.MissionName,
		Detail:      req.Detail,
		Deadline:    gtime.New(req.Deadline),
		MissionId:   req.MissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.MissionUpdateRes{MissionId: out.MissionId}, nil
}

func (a *cMission) GetListOfUser(ctx context.Context, req *backend.MissionQueryByUserIdReq) (res *backend.MissionQueryByUserIdRes, err error) {
	list, err := service.Mission().GetListOfUser(ctx, model.MissionQueryByUserIdInput{
		Page:   req.Page,
		Size:   req.Size,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.MissionQueryByUserIdRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}, nil
}
