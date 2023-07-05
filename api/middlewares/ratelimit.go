package middlewares

import (
	"log"
	"time"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type LimiterConfig struct {
	Next         func(c *fiber.Ctx) bool
	Max          int
	KeyGenerator func(*fiber.Ctx) string
	Expiration   time.Duration
	LimitReached fiber.Handler
	Storage      fiber.Storage
}

func KeyGenerator(ctx *fiber.Ctx) string {
	return ctx.Get("x-forwarded-for")
}

func LimitHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusTooManyRequests)
}

func NewRateLimiter(cfg config.RateLimiter) limiter.Config {
	log.Println(cfg)
	return limiter.Config{
		Max:               cfg.MaxReq,
		Expiration:        cfg.Expiration,
		KeyGenerator:      KeyGenerator,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached:      LimitHandler,
	}
}
