package users

import app "github.com/ZuhybDev/geeyeApp/internal"

type UserHandler struct {
	app *app.App
}

func NewUserHandler(a *app.App) *UserHandler {
	return &UserHandler{app: a}
}
