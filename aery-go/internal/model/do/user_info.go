// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta     `orm:"table:user_info, do:true"`
	UserInfoId interface{} //
	UserId     interface{} //
	Info       interface{} //
	NickName   interface{} //
	Account    interface{} //
	Sex        interface{} //
	Birth      *gtime.Time //
	Phone      interface{} //
	Brief      interface{} //
	CreateDate *gtime.Time //
	Avatar     interface{} //
	Address    interface{} //
}
