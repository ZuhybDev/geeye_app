package restaurant

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResAddressParam struct {
	StreetName *string `json:"street_name"`
	City       *string `json:"city"`
	State      *string `json:"state"`
	Phone      *string `json:"phone"`
	Email      *string `json:"email"`
	IsDefault  bool    `json:"is_default"`
}

func (h *RestaurantHandler) CreateResAddress(c fiber.Ctx) error {

	// incoming params
	var resParams ResAddressParam
	if err := c.Bind().Body(&resParams); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	resid, err := GetResId(c, h)

	if err != nil {
		fmt.Printf("DEBUG ERROR check restaurant: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Restaturant not found",
		})
	}

	params := db.CreateResAddressParams{
		RestaurantID: resid,
	}

	if resParams.StreetName != nil {
		params.StreetName = pgtype.Text{
			String: *resParams.StreetName,
			Valid:  true,
		}
	}
	if resParams.City != nil {
		params.City = pgtype.Text{
			String: *resParams.City,
			Valid:  true,
		}
	}
	if resParams.Email != nil {
		params.Email = pgtype.Text{
			String: *resParams.Email,
			Valid:  true,
		}
	}
	if resParams.Phone != nil {
		params.Phone = pgtype.Text{
			String: *resParams.Phone,
			Valid:  true,
		}
	}
	if resParams.State != nil {
		params.State = pgtype.Text{
			String: *resParams.State,
			Valid:  true,
		}
	}

	// if the user has many brach address we have update accordingly if the
	// user make newly created is_default = true, make others false else just save it

	if resParams.IsDefault {
		err := h.Cfg.Query.UpdateDefaultResBranch(c.Context(), resid)
		if err != nil {
			fmt.Printf("DEBUG ERROR return address: %v\n", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to update branches",
			})
		}
	}

	params.IsDefault = pgtype.Bool{
		Bool:  resParams.IsDefault,
		Valid: true,
	}

	address, err := h.Cfg.Query.CreateResAddress(c.Context(), params)

	if err != nil {
		fmt.Printf("DEBUG ERROR return address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to save data",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Adress successfully added",
		"address": address,
	})

}

func (h *RestaurantHandler) GetAdderessById(c fiber.Ctx) error {

	resId, err := GetResId(c, h)

	if err != nil {
		fmt.Printf("DEBUG ERROR address: %v\n", err)
		return c.Status(401).JSON(fiber.Map{
			"message": "Failed parse ID",
		})
	}

	result, err := h.Cfg.Query.GetUserResAddressesById(c.Context(), resId)

	if err != nil || len(result) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Restaurant does not have address",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":   "User addresses",
		"addresses": result,
	})

}

func (h *RestaurantHandler) UpdateResAddress(c fiber.Ctx) error {
	//copy of the RessAddressParam
	var resUpdateParams ResAddressParam

	if err := c.Bind().Body(&resUpdateParams); err != nil {
		fmt.Printf("DEBUG ERROR body in address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Invlaid request body",
		})
	}

	// params id
	paramId := c.Params("id")

	// Parse the params id into pgtype
	parsedParamId, err := utils.ParsePGIDs(paramId)

	if err != nil {
		fmt.Printf("DEBUG ERROR address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed parse ID",
		})
	}

	params := db.UpdateResAddressParams{
		ID: parsedParamId,
	}

	if resUpdateParams.StreetName != nil {
		params.StreetName = pgtype.Text{
			String: *resUpdateParams.StreetName,
			Valid:  true,
		}
	}
	if resUpdateParams.State != nil {
		params.State = pgtype.Text{
			String: *resUpdateParams.State,
			Valid:  true,
		}
	}
	if resUpdateParams.Phone != nil {
		params.Phone = pgtype.Text{
			String: *resUpdateParams.Phone,
			Valid:  true,
		}
	}

	if resUpdateParams.Email != nil {
		params.Email = pgtype.Text{
			String: *resUpdateParams.Email,
			Valid:  true,
		}
	}

	if resUpdateParams.City != nil {
		params.City = pgtype.Text{
			String: *resUpdateParams.City,
			Valid:  true,
		}
	}

	// GetResId returns current user restaurant ID
	userResId, err := GetResId(c, h)

	if resUpdateParams.IsDefault {
		err := h.Cfg.Query.UpdateDefaultResBranch(c.Context(), userResId)
		if err != nil {
			fmt.Printf("DEBUG ERROR return address: %v\n", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to update address",
			})
		}
	}

	params.IsDefault = pgtype.Bool{
		Bool:  resUpdateParams.IsDefault,
		Valid: true,
	}

	result, err := h.Cfg.Query.UpdateResAddress(c.Context(), params)

	return c.Status(200).JSON(fiber.Map{
		"message":   "Address successfully updated",
		"addresses": result,
	})
}

func (h *RestaurantHandler) DeleteAddress(c fiber.Ctx) error {

	paramId := c.Params("id")

	parsedParamId, err := utils.ParsePGIDs(paramId)

	if err != nil {
		fmt.Printf("DEBUG ERROR body in address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to parse ID",
		})
	}

	err = h.Cfg.Query.DeleteResAddress(c.Context(), parsedParamId)

	if err != nil {
		fmt.Printf("DEBUG ERROR body in address: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Address does not exist or already deleted",
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"message": "Address successfully deleted",
	})
}
