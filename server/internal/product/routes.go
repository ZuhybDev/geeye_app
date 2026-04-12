package products

import "github.com/gofiber/fiber/v3"

func RegisterProductRoutees(api fiber.Router, h *ProductsHandler) {

	api.Get("", nil)

}
