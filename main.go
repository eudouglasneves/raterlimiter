package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"myapp/limiter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %v", err)
	}

	// Configuração do Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	rateLimitIP, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	rateLimitToken, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockDuration, _ := strconv.Atoi(os.Getenv("BLOCK_DURATION"))

	lim := limiter.NewLimiter(redisClient, rateLimitIP, rateLimitToken, time.Duration(blockDuration)*time.Second)

	r := gin.Default()
	r.Use(limiter.RateLimiterMiddleware(lim))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bem-vindo ao Rate Limited API!"})
	})

	log.Fatal(r.Run(":8080"))
}
