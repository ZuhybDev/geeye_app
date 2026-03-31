package handlers

import (
	"log"
	"time"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// function get all users
func (h *Handler) GetListUsers(c fiber.Ctx) error {
	ctx := c.Context()

	users, err := h.Query.GetUserList(ctx)
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// 2. Fiber's c.JSON automatically sets the status to 200
	return c.JSON(users)
}

// expecting from user
type RegisterRequest struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	PhoneNumber  *string `json:"phone_number"`
	ImageUrl     *string `json:"image_url"`
	RestaurantID *string `json:"restaurant_id"`
}

func (h *Handler) NewUser(c fiber.Ctx) error {

	var req RegisterRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	//check email
	_, err := h.Query.CheckEmail(c.Context(), req.Email)

	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email already taken."})
	}

	// 1. Hash the password (using bcrypt)
	hashedPass, err := utils.HashedPassword([]byte(req.Password))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to process password"})
	}

	// check the email

	// 2. Prepare sqlc params (Handling the pgtype 'Valid' flag)
	params := db.NewUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPass,
	}

	// Handle optional fields
	if req.PhoneNumber != nil {
		params.PhoneNumber = pgtype.Text{String: *req.PhoneNumber, Valid: true}
	}
	if req.ImageUrl != nil {
		params.ImageUrl = pgtype.Text{String: *req.ImageUrl, Valid: true}
	}
	if req.RestaurantID != nil {
		parsedID, err := uuid.Parse(*req.RestaurantID)
		if err != nil {
			return err
		}

		params.RestaurantID = pgtype.UUID{
			Bytes: parsedID,
			Valid: true,
		}
	}

	//RestaurantID is a UUID
	_ = params.RestaurantID.Scan(req.RestaurantID)

	// 3. Save to Database
	insertUser, err := h.Query.NewUser(c.Context(), params)
	if err != nil {
		// NEVER use log.Fatal in a handler! It kills the server process.
		log.Printf("DB Error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}

	claims := utils.UserPayload{
		ID:           insertUser.ID.String(),
		Name:         insertUser.Name,
		Email:        insertUser.Email,
		RestaurentID: insertUser.RestaurantID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    "geeye-app",
		},
	}

	// 4. Generate JWT using the REAL ID from the database
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tkn, err := token.SignedString([]byte(h.JwtSecret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// 4. Set the Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,  // Important: Prevents JS from stealing the token
		Secure:   false, //TODO Set to true in production with HTTPS
		SameSite: "Lax",
	})

	// Don't send the password back!
	insertUser.Password = ""

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    insertUser,
	})
}
