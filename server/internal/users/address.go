package users

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type NewAddressParams struct {
	City      *string `json:"city" ` //validate:"required"
	State     *string `json:"state"`
	ZipCode   *string `json:"zip_code"`
	IsDefault bool    `json:"is_default"`
}

// This funcion creates new address to current user
func (h *UserHandler) NewUserAddress(c fiber.Ctx) error {

	var addressParams NewAddressParams

	if err := c.Bind().Body(&addressParams); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}

	localUser := c.Locals("user").(*middleware.UserPayload)

	userId, err := utils.ParsePGIDs(localUser.ID)

	// check if the user exist TODO

	if err != nil {
		fmt.Printf("DEBUGING: parsing userid in user address: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse user ID",
		})
	}

	params := db.CreateUserAddressParams{
		UserID:  userId,
		City:    utils.ToPgTex(addressParams.City),
		State:   utils.ToPgTex(addressParams.State),
		ZipCode: utils.ToPgTex(addressParams.ZipCode),
	}

	if addressParams.IsDefault {
		err := h.app.Query.SetDefaultUserAddress(c.Context(), userId)
		if err != nil {
			fmt.Printf("DEBUGING: db setDefault address: %v\n", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	params.IsDefault = pgtype.Bool{
		Bool:  addressParams.IsDefault,
		Valid: true,
	}

	result, err := h.app.Query.CreateUserAddress(c.Context(), params)

	if err != nil {
		fmt.Printf("DEBUGING: error returning result address from db")
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Address successfully added",
		"address": result,
	})
}

// This function modifies address
func (h *UserHandler) UpdateAddress(c fiber.Ctx) error {

	paramsId := c.Params("id")
	var updateAddress NewAddressParams

	if err := c.Bind().Body(&updateAddress); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	addressId, err := utils.ParsePGIDs(paramsId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse Address ID",
		})
	}

	userId, err := GetUserId(c, h)

	if err != nil {
		fmt.Println("parsing Id Error: ", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse user ID",
		})
	}

	res, err := h.app.Query.GetUserAddress(c.Context(), userId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Address does tot exist",
		})
	}

	for _, r := range res {
		if r.UserID != userId {
			return c.Status(401).JSON(fiber.Map{
				"message": "Anauthorized",
			})
		}
	}

	params := db.UpdateUserAddressParams{
		ID:      addressId,
		City:    utils.ToPgTex(updateAddress.City),
		State:   utils.ToPgTex(updateAddress.State),
		ZipCode: utils.ToPgTex(updateAddress.ZipCode),
	}

	if updateAddress.IsDefault {
		err := h.app.Query.SetDefaultUserAddress(c.Context(), userId)
		if err != nil {
			fmt.Printf("DEBUGING: db setDefault address: %v\n", err)
			return c.Status(400).JSON(fiber.Map{
				"message": "Failed to update user addresses",
			})
		}
	}

	params.IsDefault = pgtype.Bool{
		Bool:  updateAddress.IsDefault,
		Valid: true,
	}

	result, err := h.app.Query.UpdateUserAddress(c.Context(), params)

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully udpated user address",
		"address": result,
	})
}

// This function returns all current user addresses
func (h *UserHandler) GetUserAddresses(c fiber.Ctx) error {
	curerntUserId, err := GetUserId(c, h)

	if err != nil {
		fmt.Println("DEBUG ERROR get user id", err)
		c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse user ID",
		})
	}

	_, err = h.app.Query.GetUserById(c.Context(), curerntUserId)

	if err != nil {
		fmt.Println("DEBUG ERROR check user existance", err)
		c.Status(404).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	restul, err := h.app.Query.GetUserAddress(c.Context(), curerntUserId)

	if err != nil {
		fmt.Println("DEBUG ERROR get user addresses", err)
		c.Status(404).JSON(fiber.Map{
			"message": "User does not have any address please create one",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":   "User addresses",
		"addresses": restul,
	})

}

func (h *UserHandler) deleteUserAddress(c fiber.Ctx) error {
	paramId := c.Params("id")

	addressId, err := utils.ParsePGIDs(paramId)

	if err != nil {
		fmt.Println("DEBUG ERROR: delete user address", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse address id",
		})
	}

	// currentUserId, err := GetUserId(c, h)

	err = h.app.Query.DeleteUserAddress(c.Context(), addressId)

	if err != nil {
		fmt.Println("DEBUG ERROR: delete user address", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to delete address",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully deleted user address",
	})
}
