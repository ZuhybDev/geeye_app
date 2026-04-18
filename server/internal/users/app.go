package users

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type UserHandler struct {
	app *env.Config
}

func NewUserHandler(a *env.Config) *UserHandler {
	return &UserHandler{app: a}
}
