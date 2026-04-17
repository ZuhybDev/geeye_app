package cars

import (
	app "github.com/ZuhybDev/geeyeApp/internal"
)

type carHandler struct {
	app *app.App
}

func NewCarHandler(a *app.App) *carHandler {
	return &carHandler{app: a}
}
