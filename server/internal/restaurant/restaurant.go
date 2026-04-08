package restaurant

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RestParam struct {
	Name *string `json:"name"`
}

func (h *ResHandler) NewRestaurent(c fiber.Ctx) error {

	// jwt user data
	user := c.Locals("user").(*middleware.UserPayload)

	var restParam RestParam

	if err := c.Bind().Body(&restParam); err != nil {
		c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if restParam.Name == nil {
		c.Status(400).JSON(fiber.Map{
			"message": "Name is required.",
		})
	}

	res, err := h.app.Query.NewResTaurant(c.Context(), *restParam.Name)

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

func (h *ResHandler) UpdateRestaurant(c fiber.Ctx) error {

	var params RestParam

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

func (h *ResHandler) GetRestaurant(c fiber.Ctx) error {

	resId, err := GetResId(c, h)

	if err != nil {
		fmt.Printf("DEBUG ERROR get restaurant: %p", err)
		return c.Status(404).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	result, err := h.app.Query.GetRestaurant(c.Context(), resId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Restaurant does not exist",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"restaurant": result,
	})
}

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
