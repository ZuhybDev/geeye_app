package delivers

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type DeliverHandler struct {
	app *env.Config
}

func NewDeliverHandler(a *env.Config) *DeliverHandler {
	return &DeliverHandler{app: a}
}
