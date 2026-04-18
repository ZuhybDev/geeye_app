package users

import (
	"fmt"
	"log"
	"time"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// This function returns all users
func (h *UserHandler) GetListUsers(c fiber.Ctx) error {
	ctx := c.Context()

	users, err := h.app.Query.GetUserList(ctx)
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// 2. Fiber's c.JSON automatically sets the status to 200
	return c.JSON(users)
}

// expecting from user
type NewUserParams struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	PhoneNumber  *string `json:"phone_number"`
	ImageUrl     *string `json:"image_url"`
	RestaurantID *string `json:"restaurant_id"`
}

// this function ccrates new user
func (h *UserHandler) NewUser(c fiber.Ctx) error {

	var req NewUserParams

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	//check email
	_, err := h.app.Query.CheckEmail(c.Context(), req.Email)

	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email already taken"})
	}

	// 1. Hash the password (using bcrypt)
	hashedPass, err := utils.HashedPassword([]byte(req.Password))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to process password"})
	}

	// check the email

	// 2. Prepare sqlc params (Handling the pgtype 'Valid' flag)
	params := db.NewUserParams{
		Name:        req.Name,
		Email:       req.Email,
		Password:    hashedPass,
		PhoneNumber: utils.ToPgTex(req.PhoneNumber),
		ImageUrl:    utils.ToPgTex(req.ImageUrl),
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

	if len(params.Password) < 6 {
		return c.Status(400).JSON(fiber.Map{"message": "Password must greater then 6 charectors"})
	}

	// 3. Save to Database
	insertUser, err := h.app.Query.NewUser(c.Context(), params)
	if err != nil {
		log.Printf("DB error: %v", err)
		return c.Status(500).JSON(fiber.Map{"message": "Could not create user"})
	}

	claims := middleware.UserPayload{
		ID:    insertUser.ID.String(),
		Name:  insertUser.Name,
		Email: insertUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    "geeye-app",
		},
	}

	// 4. Generate JWT using the REAL ID from the database
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tkn, err := token.SignedString([]byte(h.app.JWTSecret))
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

// expecting from user
type UpdateUserParams struct {
	Name         *string `json:"name"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
	PhoneNumber  *string `json:"phone_number"`
	ImageUrl     *string `json:"image_url"`
	RestaurantID *string `json:"restaurant_id"`
}

// This function updates user data
func (h *UserHandler) UpdateUser(c fiber.Ctx) error {

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

// This functiosn deletes active user
func (h *UserHandler) DeleteUser(c fiber.Ctx) error {

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

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// This function allows users to login and asgin cookies
func (h *UserHandler) Login(c fiber.Ctx) error {

	var lgnUser LoginUser

	if err := c.Bind().Body(&lgnUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
	}

	if lgnUser.Email == "" || lgnUser.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Email and password required"})
	}

	res, err := h.app.Query.UserLogin(c.Context(), lgnUser.Email)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	ok := utils.VerifyPassword(lgnUser.Password, res.Password)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	claims := middleware.UserPayload{
		ID:    res.ID.String(),
		Name:  res.Name,
		Email: res.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    "geeye-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(h.app.JWTSecret)
	tkn, err := token.SignedString(secret)

	if err != nil {
		fmt.Println("DEBUG ERROR JWT Asigning", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
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

	// Don't send the password back!
	res.Password = ""

	return c.Status(200).JSON(fiber.Map{
		"message": "Welcome Back!! " + res.Name,
		"user":    res,
	})
}
