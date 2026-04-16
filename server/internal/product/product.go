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

// This function creates a new product
func (h *ProductsHandler) NewProducts(c fiber.Ctx) error {
	var productData Product
	if err := c.Bind().Body(&productData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}

	if productData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	if productData.StockQuantity <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Quantity must be greater than 0",
		})
	}

	if !productData.Price.Valid {
		return c.Status(400).JSON(fiber.Map{
			"message": "Price is required",
		})
	}

	val, err := productData.Price.Float64Value()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid price format",
		})
	}

	if val.Float64 <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Price must be greater than 0",
		})
	}

	resid, err := GetUserResId(c, h)
	if err != nil {
		fmt.Println("DEBUG ERROR products: get user resId", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid restaurant ID or does not exist",
		})
	}

	params := db.NewProductParams{
		RestaurantID:  resid,
		Name:          productData.Name,
		Description:   utils.ToPgTex(productData.Description),
		Price:         productData.Price,
		Category:      utils.ToPgTex(productData.Category),
		Images:        productData.Images,
		StockQuantity: pgtype.Int4{Int32: productData.StockQuantity, Valid: true},
	}

	product, err := h.app.Query.NewProduct(c.Context(), params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create product",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Product successfully created",
		"product": product,
	})
}

type UpdateProduct struct {
	Name          string         `json:"name"`
	Description   *string        `json:"description"`
	Price         pgtype.Numeric `json:"price"`
	Category      *string        `json:"category"`
	Images        *[]string      `json:"images"`
	StockQuantity int32          `json:"stock_quantity"`
}

// This function Updates an existing products
func (h *ProductsHandler) UpdateProduct(c fiber.Ctx) error {

	var updateProdParams UpdateProduct

	productId := c.Params("id")
	parsedPrpductId, err := utils.ParsePGIDs(productId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse product id",
		})
	}

	if err := c.Bind().Body(&updateProdParams); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}

	userResId, err := GetUserResId(c, h)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid restaurant id",
		})
	}

	products, err := h.app.Query.GetProducts(c.Context(), parsedPrpductId)

	if products.RestaurantID != userResId {
		return c.Status(402).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	params := db.UpdateProductParams{
		ID:          parsedPrpductId,
		Name:        utils.ToPgTex(&updateProdParams.Name),
		Category:    utils.ToPgTex(updateProdParams.Category),
		Description: utils.ToPgTex(updateProdParams.Description),
		Price:       updateProdParams.Price,
		Images:      *updateProdParams.Images,
		StockQuantity: pgtype.Int4{
			Int32: updateProdParams.StockQuantity,
			Valid: true,
		},
	}

	updatedProduct, err := h.app.Query.UpdateProduct(c.Context(), params)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully updated product",
		"product": updatedProduct,
	})

}

func (h *ProductsHandler) DeleteProduct(c fiber.Ctx) error {

	id := c.Params("id")

	productId, err := utils.ParsePGIDs(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse product id",
		})
	}

	userResId, err := GetUserResId(c, h)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid restaurant id",
		})
	}

	products, err := h.app.Query.GetProducts(c.Context(), productId)

	if products.RestaurantID != userResId {
		return c.Status(402).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	err = h.app.Query.DeleteProductById(c.Context(), productId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to delete product",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully deleted",
	})
}
