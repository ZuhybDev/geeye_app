package utils

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetUserResId(c fiber.Ctx, qry *db.Queries) (pgtype.UUID, error) {
	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedUserId, err := ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Printf("DEBUG ERROR return address: %v\n", err)
		return pgtype.UUID{}, err
	}

	resId, err := qry.GetUserResById(c.Context(), parsedUserId)

	return resId, nil
}

// this function returns user UUID
func GetCurrentUserId(c fiber.Ctx, h *db.Queries) (pgtype.UUID, error) {

	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedUserId, err := ParsePGIDs(localUser.ID)

	if err != nil {
		return pgtype.UUID{}, err
	}
	return parsedUserId, nil
}
