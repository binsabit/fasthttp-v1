package api

import "github.com/gofiber/fiber/v2"

func (app *Application) SetupRoutes() {

	//created product group
	product := app.Router.Group("/product")

	product.Get("", app.handlerGetProduct)
	product.Post("", app.handlerPostProduct)
	product.Get("/:id", app.handlerGetProductByID)
	product.Delete("/:id", app.handlerDeleteProduct)
	product.Put("/:id", app.handlerUpdateProduct)

	app.Router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
