package handlers

import (
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RestParam struct {
	Name string `json:"name"`
}

func (h *Handler) NewRestaurent(c fiber.Ctx) error {

	// jwt user data
	user := c.Locals("user").(*utils.UserPayload)

	var restParam RestParam

	if err := c.Bind().Body(&restParam); err != nil {
		c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if restParam.Name == "" {
		c.Status(400).JSON(fiber.Map{
			"message": "Name is required.",
		})
	}

	res, err := h.Query.NewResTaurant(c.Context(), restParam.Name)

	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	var UserParam UpdateUserParams

	parseUserId, err := uuid.Parse(user.ID)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "failed to parse id",
		})
	}

	dbId := pgtype.UUID{
		Bytes: parseUserId,
		Valid: true,
	}

	params := db.UpdateUserParams{
		ID: dbId,
	}

	parseId := uuid.UUID(res.ID.Bytes)

	if UserParam.RestaurantID != nil {
		params.RestaurantID = pgtype.UUID{
			Bytes: parseId,
			Valid: true,
		}
	}

	u, err := h.Query.UpdateUser(c.Context(), params)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":     "Restaurant successfully created",
		"restaurant":  res,
		"updatedUser": u,
	})
}
