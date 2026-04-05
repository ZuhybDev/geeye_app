package restaurant

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResAddressParam struct {
	RestaurantID string  `json:"restaurantId"`
	StreetName   *string `json:"streetName"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Phone        *string `json:"phone"`
	Email        *string `json:"email"`
	IsDefault    bool    `json:"isDefualt"`
}

func (h *ResHandler) CreateResAddress(c fiber.Ctx) error {

	localUser := c.Locals("user").(middleware.UserPayload)

	id, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Printf("DEBUG ERROR res_address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	id, err = h.app.Query.GetUserResById(c.Context(), id)

	if err != nil {
		fmt.Printf("DEBUG ERROR check restaurant: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaturant not found",
		})
	}

	// incoming params
	var resParams ResAddressParam

	params := db.CreateResAddressParams{
		RestaurantID: id,
	}

	if resParams.StreetName != nil {
		params.StreetName = pgtype.Text{
			String: *resParams.StreetName,
			Valid:  true,
		}
	}
	if resParams.City != nil {
		params.City = pgtype.Text{
			String: *resParams.StreetName,
			Valid:  true,
		}
	}
	if resParams.Email != nil {
		params.Email = pgtype.Text{
			String: *resParams.StreetName,
			Valid:  true,
		}
	}
	if resParams.Phone != nil {
		params.Phone = pgtype.Text{
			String: *resParams.StreetName,
			Valid:  true,
		}
	}

	params.IsDefault = pgtype.Bool{
		Bool:  resParams.IsDefault,
		Valid: true,
	}

	address, err := h.app.Query.CreateResAddress(c.Context(), params)

	if err != nil {
		fmt.Printf("DEBUG ERROR return address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to save data",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Adress successfully added",
		"address": address,
	})

}
