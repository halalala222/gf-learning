// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Participant is the golang structure for table participant.
type Participant struct {
	Id        int         `json:"id"        ` //
	StuId     string      `json:"stuId"     ` //
	Qq        int64       `json:"qq"        ` //
	Name      string      `json:"name"      ` //
	Question  string      `json:"question"  ` //
	Puid      int         `json:"puid"      ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	IsCheckIn int         `json:"isCheckIn" ` //
}
