package model

import "github.com/gogf/gf/v2/os/gtime"

type MissionCreateInput struct {
	Missionname string
	Detail      string
	Deadline    *gtime.Time
	UserId      int
}

type MissionCreateOutput struct {
	MissionId int `json:"id"`
}

type MissionDeleteInput struct {
	MissionId int
}

type MissionDeleteOutput struct {
	MissionId int
}

type MissionUpdateInput struct {
	MissionId   int
	Detail      string
	MissionName string
	Deadline    *gtime.Time
}

type MissionUpdateOutput struct {
	MissionId   int
	MissionName string
	Deadline    *gtime.Time
}

type MissionQueryByUserIdInput struct {
	UserId int
	Page   int `json:"page" description:"分页码"`
	Size   int `json:"size" description:"分页数量"`
}

type MissionQueryByUserIdOutput struct {
	List  []MissionQueryByUserOutputItem `json:"list" description:"列表"`
	Page  int                            `json:"page" description:"分页码"`
	Size  int                            `json:"size" description:"分页数量"`
	Total int                            `json:"total" description:"数据总数"`
}

type MissionQueryByUserOutputItem struct {
	Id         int         `json:"id"`   //
	Name       string      `json:"name"` //
	Detail     string      `json:"detail"`
	Status     string      `json:"status"` //
	CreateUid  string      `json:"userId"` //
	Deadline   *gtime.Time `json:"dealline"`
	Createtime *gtime.Time `json:"createtime"` //
	Updatetime *gtime.Time `json:"updatetime"` //
}
