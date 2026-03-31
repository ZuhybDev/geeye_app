package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) DeleteUser(c fiber.Ctx) error {

	// get the id
	// verify it
	// parse it
	// and executes it
	//

	id := c.Params("id")

	parsedId, err := uuid.Parse(id)

	dbId := pgtype.UUID{
		Bytes: parsedId,
		Valid: true,
	}

	_, err = h.Query.GetUserById(c.Context(), dbId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User does not exist.",
		})
	}

	// delete function

	err = h.Query.DeleteUser(c.Context(), dbId)

	return c.Status(200).JSON(fiber.Map{
		"message": "User successfully deleted",
	})
}
