package users

import (
	"fmt"
	"log"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateUserParams struct {
	Name         *string `json:"name"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
	PhoneNumber  *string `json:"phone_number"`
	ImageUrl     *string `json:"image_url"`
	RestaurantID *string `json:"restaurant_id"`
}

func (h *Handler) UpdateUser(c fiber.Ctx) error {

	id := c.Params("id")
	fmt.Println("id:", id)
	parsedId, err := uuid.Parse(id)

	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// parse the id
	dbId := pgtype.UUID{
		Bytes: parsedId,
		Valid: true,
	}
	_, err = h.app.Query.GetUserById(c.Context(), dbId) // check only if its exist

	if err != nil {
		log.Println("Error fetching user function:", err)
		return c.Status(404).JSON(fiber.Map{"error": "User doesnt exist."})
	}

	var UpdateParams UpdateUserParams

	if err := c.Bind().Body(&UpdateParams); err != nil {
		log.Println("Error Binding user update inputs:", err)
		return c.Status(401).JSON(fiber.Map{"error": "Internal server error"})
	}

	params := db.UpdateUserParams{
		ID: dbId,
	}

	if UpdateParams.Name != nil {
		params.Name = pgtype.Text{String: *UpdateParams.Name, Valid: true}
	}
	if UpdateParams.Email != nil {
		params.Email = pgtype.Text{String: *UpdateParams.Email, Valid: true}
	}
	if UpdateParams.Password != nil {

		bytePass := []byte(*UpdateParams.Password)
		hashedPass, err := utils.HashedPassword(bytePass)

		if err != nil {
			return err
		}
		params.Password = pgtype.Text{String: hashedPass, Valid: true}
	}
	if UpdateParams.PhoneNumber != nil {
		params.PhoneNumber = pgtype.Text{String: *UpdateParams.PhoneNumber, Valid: true}
	}
	if UpdateParams.ImageUrl != nil {
		params.ImageUrl = pgtype.Text{String: *UpdateParams.ImageUrl, Valid: true}
	}
	if UpdateParams.RestaurantID != nil {
		parsedID, err := uuid.Parse(*UpdateParams.RestaurantID)
		if err != nil {
			return err
		}
		params.RestaurantID = pgtype.UUID{Bytes: parsedID, Valid: true}
	}

	_, err = h.app.Query.UpdateUser(c.Context(), params)

	if err != nil {
		log.Println("Error Updating user:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}
