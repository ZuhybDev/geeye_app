package handlers

import "github.com/ZuhybDev/geeyeApp/db"

// helper
type QueryEnv struct {
	Query     *db.Queries
	JwtSecret string
}
