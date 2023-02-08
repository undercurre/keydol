package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MissionCreateReq struct {
	g.Meta      `path:"/backend/mission/create" tags:"Mission" method:"post" summary:"Create a new Mission"`
	MissionName string `json:"missionname" v:"required#任务名不能为空"  	dc:"任务名"`
	Detail      string `json:"detail"  	dc:"任务详情"`
	Deadline    string `json:"deadline" v:"required#最后期限不能为空"  	dc:"最后期限"`
	UserId      int    `json:"id" v:"required#id不能为空"  	dc:"id"`
}

type MissionCreateRes struct {
	MissionId int `json:"id"`
}

type MissionDeleteReq struct {
	g.Meta    `path:"/backend/mission/delete" tags:"Mission" method:"post" summary:"Delete a Mission"`
	MissionId int `json:"id" v:"required#id不能为空"  	dc:"id"`
}

type MissionDeleteRes struct {
	MissionId int `json:"id"`
}

type MissionUpdateReq struct {
	g.Meta      `path:"/backend/mission/update" tags:"Mission" method:"post" summary:"Update a Mission"`
	MissionId   int    `json:"id" v:"required#id不能为空"  	dc:"id"`
	MissionName string `json:"missionname" v:"required#任务名不能为空"  	dc:"任务名"`
	Detail      string `json:"detail"  	dc:"任务详情"`
	Deadline    string `json:"deadline" v:"required#最后期限不能为空"  	dc:"最后期限"`
}

type MissionUpdateRes struct {
	MissionId   int    `json:"id"`
	MissionName string `json:"missionname"`
	Deadline    string `json:"deadline"`
}

type MissionQueryByUserIdReq struct {
	g.Meta `path:"/backend/mission/queryByUser" tags:"Mission" method:"get" summary:"Mission of User"`
	UserId int `json:"id" v:"required#id不能为空"  	dc:"id"`
	CommonPaginationReq
}

type MissionQueryByUserIdRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
