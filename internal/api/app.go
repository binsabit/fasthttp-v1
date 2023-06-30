package api

import (
	types "github.com/binsabit/fasthttp-v1/internal/data/types"
	"github.com/binsabit/fasthttp-v1/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

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

func (app *Application) Run() error {

	app.router.Use(cors.New())
	app.setupRoutes()
	app.router.Listen(":" + app.serverPort)

	return nil
}
