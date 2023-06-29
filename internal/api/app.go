package api

import (
	"github.com/binsabit/fasthttp-v1/internal/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Application struct {
	serverPort string
	router     *fiber.App
	storage    *data.StorageInterface
}

func NewAplication(serverPort string, storage data.StorageInterface) *Application {
	return &Application{
		serverPort: serverPort,
		router:     fiber.New(),
		storage:    &storage,
	}
}

func (app *Application) ConfigureAndRun() error {

	app.router.Use(cors.New())
	app.setupRoutes()
	app.router.Listen(":" + app.serverPort)

	return nil
}
