package restaurant

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type resParams struct {
	Name *string `json:"name"`
}

func (h *ResHandler) UpdateRestaurant(c fiber.Ctx) error {

	var params resParams

	if err := c.Bind().Body(&params); err != nil {
		fmt.Println("DEGUB ERROR UDPATE RESTAURANT: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid requarest body",
		})
	}

	resId, err := GetResId(c, h)

	if err != nil {
		fmt.Println("DEGUB ERROR UDPATE RESTAURANT GET ID FROM DB: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})
	}

	id, err := h.app.Query.CheckRestaurantID(c.Context(), resId)

	if err != nil {
		fmt.Println("DEGUB ERROR UPDATE RESTAURANT: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})
	}

	dbParams := db.UpdateRestaurantParams{
		ID: id,
	}

	if params.Name != nil {
		dbParams.Name = pgtype.Text{String: *params.Name, Valid: true}
	}

	res, err := h.app.Query.UpdateRestaurant(c.Context(), dbParams)

	return c.Status(200).JSON(fiber.Map{
		"message": "Restaurant updated successfully",
		"newName": res,
	})
}
