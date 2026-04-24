package products

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type ProductsHandler struct {
	Cfg *env.Config
}

func NewProductHandler(a *env.Config) *ProductsHandler {
	return &ProductsHandler{Cfg: a}
}
