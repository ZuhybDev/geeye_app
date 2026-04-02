package handlers

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) DeleteRestaurant(c fiber.Ctx) error {

	localUser := c.Locals("user").(*middleware.UserPayload)

	userResId, err := utils.ParsePGIDs(localUser.ID)

	if err != nil {
		fmt.Println("DEGUB ERROR: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})

	}

	id, err := h.Query.GetUserResById(c.Context(), userResId)

	if err != nil {
		fmt.Println("DEGUB ERROR: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})

	}

	hasExistId, err := h.Query.CheckRestaurantID(c.Context(), id)

	if err != nil {
		fmt.Println("DEGUB ERROR delete restaurant: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})
	}

	err = h.Query.DeleteRestaurant(c.Context(), hasExistId)

	if err != nil {
		fmt.Println("DEGUB ERROR delete restaurant: ")
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete restaurant try again",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Restaurant deleted successfully",
	})
}
