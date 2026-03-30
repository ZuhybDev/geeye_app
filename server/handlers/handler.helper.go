package handlers

import "github.com/ZuhybDev/geeyeApp/db"

// helper
type Handler struct {
	Query     *db.Queries
	JwtSecret string
}
