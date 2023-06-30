package api

import (
	"context"
	"fmt"
	"time"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	types "github.com/binsabit/fasthttp-v1/internal/data/types"
	"github.com/binsabit/fasthttp-v1/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var forceShutdownTime = time.Second * 15

type Application struct {
	serverPort string
	router     *fiber.App
	storage    *types.StorageInterface
	logger     pkg.Logger
}

func NewAplication(serverPort string, storage types.StorageInterface, logger pkg.Logger) *Application {
	return &Application{
		serverPort: serverPort,
		router:     fiber.New(),
		storage:    &storage,
		logger:     logger,
	}
}

func Run(ctx context.Context) error {
	config := config.Configure()
	logger := pkg.NewLogger(config.LogFile)

	pool, err := postgesql.NewPGXPool(context.Background(), config.DB_DSN)
	if err != nil {
		return fmt.Errorf("could not establish db connetion pool: %v", err)
	}
	defer pool.Close()

	pqStorage := postgesql.NewModels(pool)

	app := NewAplication(config.ServerPort, pqStorage, logger)

	app.router.Use(cors.New())
	app.setupRoutes()

	go func() {
		if err := app.router.Listen(":" + app.serverPort); err != nil {
			app.logger.ErrorLog.Fatalf("could not server at port:%d, %v", app.serverPort, err)
		}

	}()
	app.logger.InfoLog.Printf("listening to port$ %d", app.serverPort)

	<-ctx.Done()

	app.logger.InfoLog.Println("server gracefully shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), forceShutdownTime)
	defer cancel()

	if err := app.router.ShutdownWithContext(shutdownCtx); err != nil {
		return fmt.Errorf("shutting down server: %v", err)
	}

	select {
	case <-shutdownCtx.Done():
		return fmt.Errorf("shutdown finished: %v", ctx.Err())
	}

}
