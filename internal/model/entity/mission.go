// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Mission is the golang structure for table mission.
type Mission struct {
	Id         int         `json:"id"         ` //
	Name       string      `json:"name"       ` //
	Detail     string      `json:"detail"     ` //
	Status     int         `json:"status"     ` //
	CreateUid  int         `json:"createUid"  ` //
	Deadline   *gtime.Time `json:"deadline"   ` //
	Createtime *gtime.Time `json:"createtime" ` //
	Updatetime *gtime.Time `json:"updatetime" ` //
}
