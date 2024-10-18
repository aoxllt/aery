// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta       `orm:"table:users, do:true"`
	UserId       interface{} //
	UserName     interface{} //
	UserPasswd   interface{} //
	UserRole     interface{} // 0管理员1用户
	UserStatus   interface{} // 0正常1被封
	UserEmail    interface{} //
	NewLoginTime *gtime.Time //
}
