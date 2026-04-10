package utils

import (
	"fmt"

	app "github.com/ZuhybDev/geeyeApp/internal"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type UtilHandler struct {
	app *app.App
}

func GetResId(c fiber.Ctx, h *UtilHandler) (pgtype.UUID, error) {
	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedUserId, err := ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Printf("DEBUG ERROR return address: %v\n", err)
		return pgtype.UUID{}, err
	}

	resId, err := h.app.Query.GetUserResById(c.Context(), parsedUserId)

	return resId, nil
}

func GetUserid() {}
