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
	StreetName *string `json:"street_name"`
	City       *string `json:"city"`
	State      *string `json:"state"`
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	IsDefault  bool    `json:"is_default"`
}

func (h *ResHandler) CreateResAddress(c fiber.Ctx) error {

	localUser := c.Locals("user").(*middleware.UserPayload)

	// incoming params
	var resParams ResAddressParam
	if err := c.Bind().Body(&resParams); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	id, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Printf("DEBUG ERROR res_address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	resid, err := h.app.Query.GetUserResById(c.Context(), id)

	if err != nil {
		fmt.Printf("DEBUG ERROR check restaurant: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaturant not found",
		})
	}

	params := db.CreateResAddressParams{
		RestaurantID: resid,
	}

	if resParams.StreetName != nil {
		params.StreetName = pgtype.Text{
			String: *resParams.StreetName,
			Valid:  true,
		}
	}
	if resParams.City != nil {
		params.City = pgtype.Text{
			String: *resParams.City,
			Valid:  true,
		}
	}
	if resParams.Email != nil {
		params.Email = pgtype.Text{
			String: *resParams.Email,
			Valid:  true,
		}
	}
	if resParams.Phone != nil {
		params.Phone = pgtype.Text{
			String: *resParams.Phone,
			Valid:  true,
		}
	}
	if resParams.State != nil {
		params.State = pgtype.Text{
			String: *resParams.State,
			Valid:  true,
		}
	}

	// if the user has many brach address we have update accordingly if the
	// user make newly created is_default = true, make others false else just save it

	if resParams.IsDefault {
		err := h.app.Query.UpdateDefaultResBranch(c.Context(), resid)
		if err != nil {
			fmt.Printf("DEBUG ERROR return address: %v\n", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to update branches",
			})
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

func (h *ResHandler) GetAdderessById(c fiber.Ctx) error {

	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedUserId, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Printf("DEBUG ERROR return address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed parse ID",
		})
	}

	resId, err := h.app.Query.GetUserResById(c.Context(), parsedUserId)

	result, err := h.app.Query.GetUserResAddresses(c.Context(), resId)

	return c.Status(200).JSON(fiber.Map{
		"message":   "User addresses",
		"addresses": result,
	})

}
