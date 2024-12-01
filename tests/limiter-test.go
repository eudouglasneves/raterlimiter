package tests

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"myapp/limiter"
)

func TestLimiter(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	lim := limiter.NewLimiter(client, 5, 10, 5*time.Second)

	key := "test_key"
	for i := 0; i < 5; i++ {
		if !lim.Allow(ctx, key, 5) {
			t.Errorf("Request %d blocked unexpectedly", i+1)
		}
	}

	if lim.Allow(ctx, key, 5) {
		t.Error("Exceeded limit but still allowed")
	}
}
