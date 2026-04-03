package restaurant

import (
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RestParam struct {
	Name string `json:"name"`
}

func (h *Handler) NewRestaurent(c fiber.Ctx) error {

	// jwt user data
	user := c.Locals("user").(*middleware.UserPayload)

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

	res, err := h.app.Query.NewResTaurant(c.Context(), restParam.Name)

	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Parse user id from the cookies
	parseUserId, err := uuid.Parse(user.ID)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "failed to parse id",
		})
	}

	userdbId := pgtype.UUID{
		Bytes: parseUserId,
		Valid: true,
	}

	// parse the restuarant ID
	resId := uuid.UUID(res.ID.Bytes)

	// Pass the user id and restaurant ID
	params := db.UpdateUserParams{
		ID: userdbId,
		RestaurantID: pgtype.UUID{
			Bytes: resId,
			Valid: true,
		},
	}

	_, err = h.app.Query.UpdateUser(c.Context(), params)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":    "Restaurant successfully created",
		"restaurant": res,
	})
}
