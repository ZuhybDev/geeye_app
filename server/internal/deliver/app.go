package delivers

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type DeliverHandler struct {
	Cfg *env.Config
}

func NewDeliverHandler(a *env.Config) *DeliverHandler {
	return &DeliverHandler{
		Cfg: a,
	}
}
