package reviews

import (
	app "github.com/ZuhybDev/geeyeApp/internal"
)

type RevHandler struct {
	app *app.App
}

func NewRevHandler(a *app.App) *RevHandler {
	return &RevHandler{app: a}
}
