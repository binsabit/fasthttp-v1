package api

import (
	"log"

	"github.com/binsabit/fasthttp-v1/pkg"
	"github.com/gofiber/fiber/v2"
)

type ProductRequest struct {
	Name  string `json:"productName"`
	Brand string `json:"brandName"`
}

func (app *Application) handlerGetProduct(ctx *fiber.Ctx) error {
	app.Logger.Info("hello from GetProduct")
	return nil
}
func (app *Application) handlerPostProduct(ctx *fiber.Ctx) error {

	tempProduct := &ProductRequest{}
	err := pkg.DecodeJSONBody(ctx, tempProduct)
	if err != nil {
		return err
	}
	log.Printf("%v", tempProduct)
	return nil
}
func (app *Application) handlerGetProductByID(ctx *fiber.Ctx) error {
	log.Println("hello from GetProductByID")

	return nil
}
func (app *Application) handlerDeleteProduct(ctx *fiber.Ctx) error {
	log.Println("hello from DeleteProduct")

	return nil
}
func (app *Application) handlerUpdateProduct(ctx *fiber.Ctx) error {

	log.Println("hello from UpdateProduct")
	return nil
}
