package redis

import (
	"aery-go/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterRedisService(&redis{})
}

type redis struct{}

func (r redis) HGetValue(ctx context.Context, key string, value string) (string, error) {
	//TODO implement me
	hgetvalue, err := g.Redis().Do(ctx, "HGET", key, value)
	if err != nil {
		return "", err
	}
	fmt.Println(hgetvalue.String())
	return hgetvalue.String(), nil
}
