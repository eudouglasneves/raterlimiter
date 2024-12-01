package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Limiter struct {
	client       *redis.Client
	rateLimitIP  int
	rateLimitKey int
	blockDuration time.Duration
}

func NewLimiter(client *redis.Client, rateLimitIP, rateLimitKey int, blockDuration time.Duration) *Limiter {
	return &Limiter{
		client:       client,
		rateLimitIP:  rateLimitIP,
		rateLimitKey: rateLimitKey,
		blockDuration: blockDuration,
	}
}

func (l *Limiter) Allow(ctx context.Context, key string, limit int) bool {
	// Incrementa o contador de requisições no Redis
	pipe := l.client.TxPipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Second)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false
	}

	// Verifica se o limite foi excedido
	return incr.Val() <= int64(limit)
}

func (l *Limiter) Block(ctx context.Context, key string) {
	l.client.Set(ctx, key, "blocked", l.blockDuration)
}

func (l *Limiter) IsBlocked(ctx context.Context, key string) bool {
	status, err := l.client.Get(ctx, key).Result()
	return err == nil && status == "blocked"
}
