package api

import (
	"context"
	"time"

	"github.com/binsabit/fasthttp-v1/internal/api/middlewares"
	"github.com/binsabit/fasthttp-v1/internal/config"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	"github.com/binsabit/fasthttp-v1/internal/lib/logger/sl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/exp/slog"
)

var forceShutdownTime = time.Second * 15

type Application struct {
	ServerPort string
	Router     *fiber.App
	Logger     *slog.Logger
}

func NewAplication(serverPort string, logger *slog.Logger) *Application {
	return &Application{
		ServerPort: serverPort,
		Router:     fiber.New(),
		Logger:     logger,
	}
}

func StartApp() {
	cfg := config.MustLoad()
	logger := sl.NewLogger(cfg.LogFile)
	pool, err := postgesql.NewPGXPool(context.Background(), cfg.Storage)
	if err != nil {
		logger.Error("could not establish db connetion pool", sl.Err(err))
		return
	}
	defer pool.Close()

	app := NewAplication(cfg.HTTPServer.Address, logger)
	app.SetMidlewares(logger, cfg.RateLimiter)
	app.SetupRoutes()

	app.Router.Listen(":" + app.ServerPort)
}

func (app *Application) SetMidlewares(logger *slog.Logger, cfg config.RateLimiter) {
	app.Router.Use(helmet.New())
	app.Router.Use(recover.New())
	app.Router.Use(cors.New())
	app.Router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Router.Use(middlewares.SecurityHeaders)
	app.Router.Use(limiter.New(middlewares.NewRateLimiter(cfg)))
	app.Router.Use(middlewares.NewFiberLogger(logger))

}
