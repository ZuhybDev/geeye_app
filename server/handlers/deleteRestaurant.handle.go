package handlers

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) DeleteRestaurant(c fiber.Ctx) error {

	localUser := c.Locals("user").(*middleware.UserPayload)

	resId, err := uuid.Parse(localUser.RestaurantID)

	dbResId := pgtype.UUID{
		Bytes: resId,
		Valid: true,
	}

	if err != nil {
		fmt.Println("DEGUB ERROR: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	id, err := h.Query.CheckRestaurantID(c.Context(), dbResId)

	if err != nil {
		fmt.Println("DEGUB ERROR delete restaurant: ", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaurant not found",
		})
	}

	err = h.Query.DeleteRestaurant(c.Context(), id)

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
