package delivers

import (
	"fmt"
	"time"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ReqDeliverParams struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	LicenseNumber string `json:"license_number"`
	NationalID    string `json:"national_id"`
	CarID         string `json:"car_id"`
	SiOnline      bool   `json:"si_online"`
}

func (h *DeliverHandler) NewDeliver(c fiber.Ctx) error {

	var reqParams ReqDeliverParams

	if err := c.Bind().Body(&reqParams); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// fmt.Println(reqParams) //todo

	if len(reqParams.Password) < 8 || reqParams.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password must be greater then or equal 8",
		})
	}

	if reqParams.NationalID == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "National ID is required",
		})
	}
	if reqParams.LicenseNumber == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "license number is required",
		})
	}

	// hash the password
	hashedPassword, err := utils.HashedPassword([]byte(reqParams.Password))

	if err != nil {
		fmt.Printf("Failed to hash password: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	parsedCarId, err := utils.ParsePGIDs(reqParams.CarID)

	if err != nil {
		fmt.Printf("Failed to parse car id: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid car id",
		})
	}

	/// pass the deliver data
	params := db.NewDeliverParams{
		Name:          reqParams.Name,
		Email:         reqParams.Email,
		Password:      hashedPassword,
		LicenseNumber: utils.ToPgTex(&reqParams.LicenseNumber),
		NationalID:    utils.ToPgTex(&reqParams.NationalID),
		CarID: pgtype.UUID{
			Bytes: parsedCarId.Bytes,
			Valid: true,
		},
		SiOnline: pgtype.Bool{
			Bool:  reqParams.SiOnline,
			Valid: true,
		},
	}

	// fmt.Println(params) //todo

	deliver, err := h.Cfg.Query.NewDeliver(c.Context(), params)

	if err != nil {
		fmt.Printf("Failed to save deliver: %v\n", err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to save data deliver",
		})
	}

	claims := middleware.UserPayload{
		ID:    deliver.ID.String(),
		Name:  deliver.Name,
		Email: deliver.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    "geeye-app",
		},
	}

	// 4. Generate JWT using the REAL ID from the database
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tkn, err := token.SignedString([]byte(h.Cfg.DeliverJwtSecret))
	if err != nil {
		fmt.Printf("newDeliver: generating tokens failed")
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// 4. Set the Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,
		Secure:   false, //TODO Set to true in production with HTTPS
		SameSite: "Lax",
	})

	return c.Status(201).JSON(fiber.Map{
		"message": "Successfully created driver",
		"deliver": deliver,
	})

}
