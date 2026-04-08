package restaurant

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func (h *ResHandler) DeleteRestaurant(c fiber.Ctx) error {

	userResId, err := GetResId(c, h)

	if err != nil {
		fmt.Println("DEGUB ERROR: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})

	}

	id, err := h.app.Query.GetUserResById(c.Context(), userResId)

	if err != nil {
		fmt.Println("DEGUB ERROR: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})

	}

	hasExistId, err := h.app.Query.CheckRestaurantID(c.Context(), id)

	if err != nil {
		fmt.Println("DEGUB ERROR delete restaurant: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})
	}

	err = h.app.Query.DeleteRestaurant(c.Context(), hasExistId)

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
