package rds

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	redis *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	return &RedisClient{
		redis: redis.NewClient(&redis.Options{Addr: addr}),
	}
}

func (r *RedisClient) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	data, _ := json.Marshal(val)
	return r.redis.Set(ctx, key, data, ttl).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string, dest any) error {
	data, err := r.redis.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

func (r *RedisClient) Enque(ctx context.Context, queue string, val any) error {
	data, _ := json.Marshal(val)
	return r.redis.LPush(ctx, queue, data).Err()
}

func (r *RedisClient) Dequeue(ctx context.Context, queue string, dest any) error {
	data, err := r.redis.BRPop(ctx, 0, queue).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data[1]), dest)
}
