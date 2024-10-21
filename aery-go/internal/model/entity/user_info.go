// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	UserInfoId int         `json:"userInfoId" orm:"user_info_id" ` //
	UserId     int         `json:"userId"     orm:"user_id"      ` //
	Info       string      `json:"info"       orm:"info"         ` //
	NickName   string      `json:"nickName"   orm:"nick_name"    ` //
	Account    string      `json:"account"    orm:"account"      ` //
	Sex        string      `json:"sex"        orm:"sex"          ` //
	Birth      *gtime.Time `json:"birth"      orm:"birth"        ` //
	Phone      string      `json:"phone"      orm:"phone"        ` //
	Brief      string      `json:"brief"      orm:"brief"        ` //
	CreateDate *gtime.Time `json:"createDate" orm:"create_date"  ` //
	Avatar     string      `json:"avatar"     orm:"avatar"       ` //
	Address    string      `json:"address"    orm:"address"      ` //
}
