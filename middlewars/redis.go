package middlewares

import (
	"context"
	"fmt"
	"time"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(timeOut time.Duration, maxRequests int64) fiber.Handler {
	return func(c *fiber.Ctx) error {

		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		clientIP := c.Get("X-Forwarded-For")
		if clientIP == "" {
			clientIP = c.IP()
		}
		fmt.Println("Client IP: ", clientIP)

		requestCountKey := clientIP + "_requests"
		requestCount := client.Incr(context.Background(), requestCountKey).Val()
		fmt.Println("Request count: ", requestCount)

		if requestCount == 1 {
			client.Expire(context.Background(), requestCountKey, timeOut*time.Second)
		}

		if requestCount > maxRequests {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"status":  fiber.StatusTooManyRequests,
				"message": "You have exceeded the limit of requests",
			})
		}

		return c.Next()
	}
}
