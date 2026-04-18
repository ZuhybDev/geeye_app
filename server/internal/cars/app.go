package cars

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type carHandler struct {
	app *env.Config
}

func NewCarHandler(a *env.Config) *carHandler {
	return &carHandler{app: a}
}
