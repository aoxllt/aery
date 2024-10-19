package cookie

import "github.com/gogf/gf/v2/frame/g"

type GetCookieReq struct {
	g.Meta    `path:"/getcookie" method:"post" tags:"cookie值获取"`
	CookieKey string `json:"cookie_key" v:"required"`
}
type GetCookieRes struct {
	CookieValue string `json:"cookievalue"`
}
