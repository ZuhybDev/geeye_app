package products

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type Product struct {
	Name          string    `json:"name"`
	Description   *string   `json:"description"`
	Price         float64   `json:"price"`
	Category      *string   `json:"category"`
	Images        *[]string `json:"images"`
	StockQuantity *string   `json:"stock_quantity"`
	AverageRating *string   `json:"average_rating"`
	TotalReviews  *string   `json:"total_reviews"`
}

func (h *ProductsHandler) NewProducts(c fiber.Ctx) error {

	var productData Product

	if err := c.Bind().Body(&productData); err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}

	resid, err := GetUserResId(c, h)

	if err != nil {
		fmt.Println("DEBUG ERROR products: get user resId", err)
	return	c.Status(400).JSON(fiber.Map{
			"message": "Invalid restaurant ID or does not exist",
		})
	}

	if productData.Name == "" {
	return	c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	if productData.Price > 0 {
	return	c.Status(400).JSON(fiber.Map{
			"message": "Price must be greater than 0",
		})
	}

	// params := db.Product{
	// 	RestaurantID: resid,
	// }

	return c.Status(400).JSON(fiber.Map{
		"message": "Price must be greater than 0",
	})
}
