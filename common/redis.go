package common

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedisClient(addr string, password string, db int) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:             addr,
		Password:         password,
		DB:               db,
		DisableIndentity: true,
	})
	redisClient := RedisClient{
		client: client,
	}
	return redisClient
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) Set(key string, value any, expiration time.Duration) {
	r.client.Set(ctx, key, value, expiration)
}

func (r *RedisClient) Get(key string) string {
	return r.client.Get(ctx, key).Val()
}

func (r *RedisClient) Del(key string) bool {
	intCmd := r.client.Del(ctx, key)
	return intCmd.Val() == 1
}

func (r *RedisClient) Exist(key string) bool {
	intCmd := r.client.Exists(ctx, key)
	return intCmd.Val() == 1
}

func (r *RedisClient) Expire(key string, expiration time.Duration) bool {
	boolCmd := r.client.Expire(ctx, key, expiration)
	return boolCmd.Val()
}

func (r *RedisClient) ExpireTime(key string) *redis.DurationCmd {
	return r.client.ExpireTime(ctx, key)
}
