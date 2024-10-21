package main

import (
	"aery-go/internal/cmd"
	_ "aery-go/internal/logic"
	_ "aery-go/internal/packed"
	"aery-go/utility"
	_ "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	_ "github.com/go-redis/redis/v8"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入MySQL驱动
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/mojocn/base64Captcha"
)

func main() {
	utility.InitOssBucket()
	//fmt.Println(utility.Oss.Getsts("test/ss.jpg", "ss.jpg"))
	cmd.Main.Run(gctx.GetInitCtx())
}
