package api

import (
	"context"
	"log"
	"time"

	"github.com/binsabit/fasthttp-v1/internal/config"
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

func Run() {
	cfg := config.MustLoad()

	infoLog := pkg.NewLogger(cfg.LogFile, "INFO")
	errorLog := pkg.NewLogger(cfg.LogFile, "ERROR")

	pool, err := postgesql.NewPGXPool(context.Background(), cfg.Storage)
	if err != nil {
		errorLog.Fatalf("could not establish db connetion pool: %v", err)
	}
	defer pool.Close()
	app := NewAplication(cfg.HTTPServer.Address, infoLog, errorLog)

	app.router.Use(cors.New())
	app.setupRoutes()

	app.router.Listen(":" + app.serverPort)

}
