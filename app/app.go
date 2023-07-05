package app

import (
	"context"
	"time"

	"github.com/binsabit/fasthttp-v1/api/middlewares"
	"github.com/binsabit/fasthttp-v1/api/routes"
	"github.com/binsabit/fasthttp-v1/config"
	"github.com/binsabit/fasthttp-v1/database/postgesql"
	"github.com/binsabit/fasthttp-v1/lib/logger/sl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/exp/slog"
)

var forceShutdownTime = time.Second * 15

func StartApp() {
	cfg := config.MustLoad()
	logger := sl.NewLogger(cfg.LogFile)
	pool, err := postgesql.NewPGXPool(context.Background(), cfg.Storage)
	if err != nil {
		logger.Error("could not establish db connetion pool", sl.Err(err))
		return
	}
	defer pool.Close()
	app := fiber.New()

	RegisterMidlewares(app, logger, cfg.RateLimiter)

	routes.SetupProductRoutes(app)

	app.Listen(":" + cfg.HTTPServer.Address)
}

func RegisterMidlewares(app *fiber.App, logger *slog.Logger, cfg config.RateLimiter) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(limiter.New(middlewares.NewRateLimiter(cfg)))
	app.Use(middlewares.NewFiberLogger(logger))
}
