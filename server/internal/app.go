package internal

import "github.com/ZuhybDev/geeyeApp/db"

type App struct {
	Query     *db.Queries
	JwtSecret string
}
