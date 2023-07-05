package products

import (
	"github.com/binsabit/fasthttp-v1/lib/json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type ProductRequest struct {
	Name  string `json:"productName"`
	Brand string `json:"brandName"`
}

func HandlerGetProduct(ctx *fiber.Ctx) error {
	slog.Info("hello from GetProduct")
	return nil
}
func HandlerPostProduct(ctx *fiber.Ctx) error {

	tempProduct := &ProductRequest{}
	err := json.DecodeJSONBody(ctx, tempProduct)
	if err != nil {
		return err
	}
	slog.Info(tempProduct.Brand, tempProduct.Name)
	return nil
}
