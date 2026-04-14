package products

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetUserResId(c fiber.Ctx, h *ProductsHandler) (pgtype.UUID, error) {
	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedUserId, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Printf("DEBUG ERROR return address: %v\n", err)
		return pgtype.UUID{}, err
	}

	resId, err := h.app.Query.GetUserResById(c.Context(), parsedUserId)

	return resId, nil
}

// this function returns user UUID
func GetCurrentUserId(c fiber.Ctx, h *ProductsHandler) (pgtype.UUID, error) {

	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedUserId, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		return pgtype.UUID{}, err
	}
	return parsedUserId, nil
}
