package restaurant

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
)

type RestaurantHandler struct {
	Cfg *env.Config
}

// Accept the pointer to the entire config
func NewRestaurantHandler(config *env.Config) *RestaurantHandler {
	return &RestaurantHandler{
		Cfg: config,
	}
}
