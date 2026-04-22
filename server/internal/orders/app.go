package orders

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type OrderHandler struct {
	Cfg *env.Config
}

func NewOrderHandler(a *env.Config) *OrderHandler {
	return &OrderHandler{
		Cfg: a,
	}
}
