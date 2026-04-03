package users

import app "github.com/ZuhybDev/geeyeApp/internal"

type Handler struct {
	app *app.App
}

func NewUserHandler(a *app.App) *Handler {
	return &Handler{app: a}
}
