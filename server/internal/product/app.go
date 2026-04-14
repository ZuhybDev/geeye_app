package products

import app "github.com/ZuhybDev/geeyeApp/internal"

type ProductsHandler struct {
	app *app.App
}

func NewProductHandler(a *app.App) *ProductsHandler {
	return &ProductsHandler{app: a}
}
