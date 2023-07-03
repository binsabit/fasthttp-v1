package api

import (
	"context"
	"time"

	"github.com/binsabit/fasthttp-v1/internal/config"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	"github.com/binsabit/fasthttp-v1/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/exp/slog"
)

var forceShutdownTime = time.Second * 15

type Application struct {
	serverPort string
	router     *fiber.App
	Logger     *slog.Logger
}

func NewAplication(serverPort string, logger *slog.Logger) *Application {
	return &Application{
		serverPort: serverPort,
		router:     fiber.New(),
		Logger:     logger,
	}
}

func Run() {
	cfg := config.MustLoad()
	logger := pkg.NewLogger(cfg.LogFile)
	logger.Info(cfg.LogFile)
	pool, err := postgesql.NewPGXPool(context.Background(), cfg.Storage)
	if err != nil {
		logger.Error("could not establish db connetion pool: %v", err)
		return
	}
	defer pool.Close()

	app := NewAplication(cfg.HTTPServer.Address, logger)

	app.router.Use(cors.New())
	app.setupRoutes()

	app.router.Listen(":" + app.serverPort)

}
