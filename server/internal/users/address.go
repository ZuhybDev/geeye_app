package users

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type NewAddressParams struct {
	City      *string `json:"city" ` //validate:"required"
	State     *string `json:"state"`
	ZipCode   *string `json:"zip_code"`
	IsDefault bool    `json:"is_default"`
}

func (h *Handler) NewUserAddress(c fiber.Ctx) error {

	var addressParams NewAddressParams

	if err := c.Bind().Body(&addressParams); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}

	// if err := validate.Struct(addressParams); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"message": "validation error",
	// 	})
	// }

	localUser := c.Locals("user").(*middleware.UserPayload)

	userId, err := utils.ParsePGIDs(localUser.ID)

	// check if the user exist TODO

	if err != nil {
		fmt.Printf("DEBUGING: parsing userid in user address: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse user ID",
		})
	}

	params := db.CreateUserAddressParams{
		UserID:  userId,
		City:    ToPgTex(addressParams.City),
		State:   ToPgTex(addressParams.State),
		ZipCode: ToPgTex(addressParams.ZipCode),
	}

	if addressParams.IsDefault {
		err := h.app.Query.SetDefaultUserAddress(c.Context(), userId)
		if err != nil {
			fmt.Printf("DEBUGING: db setDefault address: %v\n", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	params.IsDefault = pgtype.Bool{
		Bool:  addressParams.IsDefault,
		Valid: true,
	}

	result, err := h.app.Query.CreateUserAddress(c.Context(), params)

	if err != nil {
		fmt.Printf("DEBUGING: error returning result address from db")
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Address successfully added",
		"address": result,
	})
}
