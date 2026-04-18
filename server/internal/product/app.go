package products

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type ProductsHandler struct {
	app *env.Config
}

func NewProductHandler(a *env.Config) *ProductsHandler {
	return &ProductsHandler{app: a}
}
