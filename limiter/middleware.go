package limiter

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(limiter *Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")

		key := ip
		limit := limiter.rateLimitIP

		if token != "" {
			key = "token:" + token
			limit = limiter.rateLimitKey
		}

		if limiter.IsBlocked(ctx, key) {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "you have reached the maximum number of requests or actions allowed within a certain time frame"})
			c.Abort()
			return
		}

		if !limiter.Allow(ctx, key, limit) {
			limiter.Block(ctx, key)
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "you have reached the maximum number of requests or actions allowed within a certain time frame"})
			c.Abort()
			return
		}

		c.Next()
	}
}
