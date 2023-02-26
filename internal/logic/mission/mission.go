package mission

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

type sMission struct{}

func init() {
	service.RegisterMission(New())
}

func New() *sMission {
	return &sMission{}
}

func (s *sMission) Create(ctx context.Context, in model.MissionCreateInput) (out model.MissionCreateOutput, err error) {
	//插入数据返回id
	lastInsertID, err := dao.Mission.Ctx(ctx).Data(g.Map{
		"name":       in.Missionname,
		"detail":     in.Detail,
		"status":     0,
		"createUid":  in.UserId,
		"deadline":   in.Deadline,
		"createTime": gdb.Raw("now()"),
		"updateTime": gdb.Raw("now()"),
	}).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.MissionCreateOutput{MissionId: int(lastInsertID)}, err
}

func (s *sMission) Delete(ctx context.Context, in model.MissionDeleteInput) (out model.MissionDeleteOutput, err error) {
	//插入数据返回id
	_, err = dao.Mission.Ctx(ctx).Where("id", in.MissionId).Delete()
	if err != nil {
		return out, err
	}
	return model.MissionDeleteOutput{MissionId: in.MissionId}, err
}

func (s *sMission) Update(ctx context.Context, in model.MissionUpdateInput) (out model.MissionUpdateOutput, err error) {
	//插入数据返回id
	_, err = dao.Mission.Ctx(ctx).Data(g.Map{
		"name":     in.MissionName,
		"status":   in.Status,
		"detail":   in.Detail,
		"deadline": in.Deadline,
	}).Where("id", in.MissionId).Update()
	if err != nil {
		return out, err
	}
	return model.MissionUpdateOutput{MissionId: in.MissionId}, err
}

func (s *sMission) GetListOfUser(ctx context.Context, in model.MissionQueryByUserIdInput) (out *model.MissionQueryByUserIdOutput, err error) {
	var (
		m = dao.Mission.Ctx(ctx)
	)
	out = &model.MissionQueryByUserIdOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size).Where("createUid=?", in.UserId)
	// 排序方式
	// listModel = listModel.OrderDesc(dao.User.Columns().Username)

	// 执行查询
	var list []*entity.Mission
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	var (
		temp, tempErr = listModel.Count()
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
