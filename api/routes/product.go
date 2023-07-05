package routes

import (
	products "github.com/binsabit/fasthttp-v1/api/handlers/product"
	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App) {

	//created product group
	product := app.Group("/product")

	product.Get("", products.HandlerGetProduct)
	product.Post("", products.HandlerPostProduct)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
