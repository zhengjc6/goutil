package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// SetClient 创建客户端
func (l *RedisQueue) SetClient(redisClient *redis.Client) *RedisQueue {
	l.client = redisClient
	return l
}

// SetKey 设置key
func (l *RedisQueue) SetKey(key string) *RedisQueue {
	l.key = key
	return l
}

// Push 推送消息
func (l *RedisQueue) Push(ctx context.Context, itm interface{}) error {
	redisResult := l.client.LPush(ctx, l.key, itm)
	if redisResult.Err() != nil {
		return redisResult.Err()
	}
	return redisResult.Err()
}

// Get 查看消息
func (l *RedisQueue) Get(ctx context.Context, start int64, stop int64) ([]string, error) {
	redisResult := l.client.LRange(ctx, l.key, start, stop)
	if redisResult.Err() != nil {
		return nil, redisResult.Err()
	}
	return redisResult.Result()
}

// Pop 取数据
func (l *RedisQueue) Pop(ctx context.Context) ([]byte, error) {
	redisResult := l.client.LPop(ctx, l.key)
	if redisResult.Err() != nil {
		return nil, redisResult.Err()
	}
	return redisResult.Bytes()
}

// Len 数量
func (l *RedisQueue) Len(ctx context.Context) (error, int64) {
	count, err := l.client.LLen(ctx, l.key).Result()
	if err != nil {
		return err, count
	}
	return nil, count
}
