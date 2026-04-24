package users

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type UserHandler struct {
	Cfg *env.Config
}

func NewUserHandler(a *env.Config) *UserHandler {
	return &UserHandler{Cfg: a}
}
