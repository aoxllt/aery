// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	UserId       int         `json:"userId"       orm:"user_id"        ` //
	UserName     string      `json:"userName"     orm:"user_name"      ` //
	UserPasswd   string      `json:"userPasswd"   orm:"user_passwd"    ` //
	UserRole     string      `json:"userRole"     orm:"user_role"      ` // 0管理员1用户
	UserStatus   string      `json:"userStatus"   orm:"user_status"    ` // 0正常1被封
	UserPhone    string      `json:"userPhone"    orm:"user_phone"     ` //
	NewLoginTime *gtime.Time `json:"newLoginTime" orm:"new_login_time" ` //
}
