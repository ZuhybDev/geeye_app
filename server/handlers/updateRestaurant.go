package handlers

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type resParams struct {
	Name *string `json:"name"`
}

func (h *QueryEnv) UpdateRestaurant(c fiber.Ctx) error {

	var params resParams

	if err := c.Bind().Body(&params); err != nil {
		fmt.Println("DEGUB ERROR UDPATE RESTAURANT: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid requarest body",
		})
	}

	localUser := c.Locals("user").(*middleware.UserPayload)

	parsedId, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Println("DEGUB ERROR UDPATE RESTAURANT PARSE ID: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	resId, err := h.Query.GetUserResById(c.Context(), parsedId)

	if err != nil {
		fmt.Println("DEGUB ERROR UDPATE RESTAURANT GET ID FROM DB: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})
	}

	id, err := h.Query.CheckRestaurantID(c.Context(), resId)

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

	res, err := h.Query.UpdateRestaurant(c.Context(), dbParams)

	return c.Status(200).JSON(fiber.Map{
		"message": "Restaurant updated successfully",
		"newName": res,
	})
}
