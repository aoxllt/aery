package service

import "context"

type IRedisService interface {
	HGetValue(ctx context.Context, key string, value string) (string, error)
}

var localredis IRedisService

func RedisServiceforController() IRedisService {
	if localredis == nil {
		panic("redis service not initialized")
	}
	return localredis
}

func RegisterRedisService(s IRedisService) {
	localredis = s
}
