package queue

import "github.com/go-redis/redis/v8"
import "context"

// Queue 接口定义队列的操作
type Queue interface {
	Push(ctx context.Context, itm interface{}) error
	Pop(ctx context.Context) ([]byte, error)
	Get(ctx context.Context, start int64, stop int64) ([]string, error)
	Len(ctx context.Context) (error, int64)
}

// RedisQueue redis的队列
type RedisQueue struct {
	client *redis.Client
	key    string
}
