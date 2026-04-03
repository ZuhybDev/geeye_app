package restaurant

import app "github.com/ZuhybDev/geeyeApp/internal"

type Handler struct {
	app *app.App
}

func NewRestaurantHandler(a *app.App) *Handler {
	return &Handler{app: a}
}
