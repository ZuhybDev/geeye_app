package users

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) DeleteUser(c fiber.Ctx) error {

	localUser := c.Locals("user").(middleware.UserPayload)

	parsedId, err := uuid.Parse(localUser.ID)

	dbId := pgtype.UUID{
		Bytes: parsedId,
		Valid: true,
	}

	_, err = h.app.Query.GetUserById(c.Context(), dbId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User does not exist.",
		})
	}

	// delete function

	err = h.app.Query.DeleteUser(c.Context(), dbId)

	return c.Status(200).JSON(fiber.Map{
		"message": "User successfully deleted",
	})
}
