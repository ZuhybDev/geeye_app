package products

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	Name          string         `json:"name"`
	Description   *string        `json:"description"`
	Price         pgtype.Numeric `json:"price"`
	Category      *string        `json:"category"`
	Images        []string       `json:"images"`
	StockQuantity int32          `json:"stock_quantity"`
}

func (h *ProductsHandler) NewProducts(c fiber.Ctx) error {

	var productData Product

	if err := c.Bind().Body(&productData); err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}

	// validation inputs
	if productData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}
	if productData.StockQuantity <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Quantity must greater then 0",
		})
	}

	if !productData.Price.Valid {
		return fmt.Errorf("price is required")
	}

	val, err := productData.Price.Float64Value()
	if err != nil {
		return fmt.Errorf("invalid price format")
	}

	if val.Float64 <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	resid, err := GetUserResId(c, h)

	if err != nil {
		fmt.Println("DEBUG ERROR products: get user resId", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid restaurant ID or does not exist",
		})
	}

	params := db.NewProductParams{
		RestaurantID: resid,
		Name:         productData.Name,
		Description:  utils.ToPgTex(productData.Description),
		Price:        productData.Price,
		Category:     utils.ToPgTex(productData.Category),
		Images:       productData.Images,
		StockQuantity: pgtype.Int4{
			Int32: productData.StockQuantity,
			Valid: true,
		},
	}

	product, err := h.app.Query.NewProduct(c.Context(), params)

	return c.Status(201).JSON(fiber.Map{
		"message": "Product successfully created",
		"product": product,
	})
}
