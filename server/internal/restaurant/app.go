package restaurant

import app "github.com/ZuhybDev/geeyeApp/internal"

type ResHandler struct {
	app *app.App
}

func NewRestaurantHandler(a *app.App) *ResHandler {
	return &ResHandler{app: a}
}
