package restaurant

type ResAddressParam struct {
	RestaurantID string `json:"restaurantId"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	State        string `json:"state"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	IsDefault    bool   `json:"isDefualt"`
}

// func (h *ResHandler) CreateResAddress(c fiber.Ctx) error {

// 	localUser := c.Locals("user").(middleware.UserPayload)

// 	resId := localUser.RestaurantID

// 	// incoming params
// 	var resParams ResAddressParam

// }
