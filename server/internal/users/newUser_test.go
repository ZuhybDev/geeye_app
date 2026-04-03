package users

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

type DBQueries interface {
	NewUser(ctx context.Context, arg db.NewUserParams) (db.User, error)
}

// ✅ Handler struct
type Handle struct {
	Query DBQueries
}

// ✅ Mock DB
type dbMock struct{}

func (m *dbMock) NewUser(ctx context.Context, arg db.NewUserParams) (db.User, error) {
	return db.User{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		Name:  arg.Name,
		Email: arg.Email,
	}, nil
}

// ✅ Handler function
func (h *Handle) NewUser(c fiber.Ctx) error {
	var req db.User

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	user, err := h.Query.NewUser(context.Background(), db.NewUserParams{
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

// ✅ TEST For this
func TestNewUser(t *testing.T) {
	mockDB := &dbMock{}
	h := &Handle{Query: mockDB}

	app := fiber.New()
	app.Post("/register", h.NewUser)

	// request body
	payload := db.User{
		Name:     "Zuhaib",
		Email:    "test@example.com",
		Password: "password123",
	}

	bodyBytes, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/register", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// run request
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	// read response
	body, _ := io.ReadAll(resp.Body)

	assert.Contains(t, string(body), "User created successfully")
	assert.Contains(t, string(body), "test@example.com")
}
