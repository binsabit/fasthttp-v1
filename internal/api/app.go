package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	"github.com/binsabit/fasthttp-v1/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var forceShutdownTime = time.Second * 15

type Application struct {
	serverPort string
	router     *fiber.App
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func NewAplication(serverPort string, infolog, errorlog *log.Logger) *Application {
	return &Application{
		serverPort: serverPort,
		router:     fiber.New(),
		InfoLog:    infolog,
		ErrorLog:   errorlog,
	}
}

func Run(ctx context.Context) error {
	config := config.Configure()

	infoLog := pkg.NewLogger(config.LogFile, "INFO")
	errorLog := pkg.NewLogger(config.LogFile, "ERROR")

	pool, err := postgesql.NewPGXPool(context.Background(), config.DB_DSN)
	if err != nil {
		return fmt.Errorf("could not establish db connetion pool: %v", err)
	}
	defer pool.Close()

	app := NewAplication(config.ServerPort, infoLog, errorLog)

	app.router.Use(cors.New())
	app.setupRoutes()

	return app.router.Listen(":" + app.serverPort)

}
