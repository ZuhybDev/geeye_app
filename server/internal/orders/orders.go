package orders

import (
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderRequest struct {
	RestaurantID    string `json:"restaurant_id"`
	PickupLocation  string `json:"pickup_location"`
	DropoffLocation string `json:"dropoff_location"`
	Items           []struct {
		ProductID string `json:"product_id"`
		Quantity  int32  `json:"quantity"`
	}
}

type ProductInfo struct {
	Price  pgtype.Numeric
	Images string
}

func (h *OrderHandler) CreatOrder(c fiber.Ctx) error {

	var orderParams OrderRequest

	if err := c.Bind().Body(&orderParams); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// combine all UUIDs one trip and grap all uu ids and related data
	var productIds []pgtype.UUID
	for _, i := range orderParams.Items {
		parsedId, _ := utils.ParsePGIDs(i.ProductID)
		productIds = append(productIds, parsedId)
	}

	productsFromDB, err := h.Cfg.Query.GetProductsByIDs(c.Context(), productIds)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error fetching prices",
		})
	}

	// I use map for quick look up
	priceMap := make(map[pgtype.UUID]ProductInfo)
	for _, p := range productsFromDB {
		priceMap[p.ID] = ProductInfo{
			Price:  p.Price,
			Images: p.Images,
		}
	}

	var items []db.CreateOrderItemParams
	var runningPrice float64

	// Order items
	for _, reqItem := range orderParams.Items {
		id, _ := utils.ParsePGIDs(reqItem.ProductID)
		prodInfo, exists := priceMap[id]

		if !exists {
			return c.Status(404).JSON(fiber.Map{"message": "One or more products not found"})
		}

		items = append(items, db.CreateOrderItemParams{
			ProductID:       id,
			Quantity:        reqItem.Quantity,
			PriceAtPurchase: prodInfo.Price,
			Image: pgtype.Text{
				String: prodInfo.Images,
				Valid:  true,
			},
		})

		priceVal, _ := prodInfo.Price.Float64Value()
		runningPrice += priceVal.Float64 * float64(reqItem.Quantity)
	}

	userId, err := utils.GetCurrentUserId(c, h.Cfg.Query)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	parsedResId, err := utils.ParsePGIDs(orderParams.RestaurantID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid restaurant ID",
		})
	}
	// Convert back to numeric
	var finalTotal pgtype.Numeric
	finalTotal.Scan(fmt.Sprintf("%f", runningPrice))

	err = h.CreateOrderTx(c.Context(), db.CreateOrderParams{
		UserID:          userId,
		RestaurantID:    parsedResId,
		TotalPrice:      finalTotal,
		PickupLocation:  utils.ToPgTex(&orderParams.PickupLocation),
		DropoffLocation: utils.ToPgTex(&orderParams.DropoffLocation),
		Status: pgtype.Text{
			String: "Pending",
			Valid:  true,
		},
	}, items)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Order placed successfully",
	})
}

type UpdateOrderRequest struct {
	PickupLocation  *string `json:"pickup_location"`
	DropoffLocation *string `json:"dropoff_location"`
	Status          *string `json:"status"`
}

func (h *OrderHandler) UpdateOrder(c fiber.Ctx) error {
	var reqParams UpdateOrderRequest

	if err := c.Bind().Body(&reqParams); err != nil {
		return c.Status(400).JSON("Invalid body request")
	}
	id, err := utils.ParsePGIDs(c.Params("id"))

	if err != nil {
		fmt.Printf("Error search params: %s\n", err)
		return c.Status(400).JSON("Internal server error")
	}

	params := db.UpdateOrderParams{
		ID:              id,
		PickupLocation:  utils.ToPgTex(reqParams.PickupLocation),
		DropoffLocation: utils.ToPgTex(reqParams.DropoffLocation),
		Status:          utils.ToPgTex(reqParams.Status),
	}

	order, err := h.Cfg.Query.UpdateOrder(c.Context(), params)

	if err != nil {
		fmt.Printf("Error updating order: %s\n", err)
		return c.Status(400).JSON("failed to update order try again")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Order Successfully Udpated",
		"order":   order,
	})

}

func (h *OrderHandler) DeleteOrder(c fiber.Ctx) error {
	id, err := utils.ParsePGIDs(c.Params("id"))

	if err != nil {
		fmt.Printf("Error search params: %s\n", err)
		return c.Status(400).JSON("Internal server error")
	}

	err = h.Cfg.Query.DeleteOrder(c.Context(), id)

	if err != nil {
		fmt.Printf("Error deleting order: %s\n", err)
		return c.Status(400).JSON("Failed to delete order try again")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Order successfully deleted",
	})
}

type DeleteParams struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

func (h *OrderHandler) DeleteOrderItems(c fiber.Ctx) error {

	var deleteReq DeleteParams

	if err := c.Bind().Body(&deleteReq); err != nil {
		return c.Status(400).JSON("invalid request body")
	}

	id, err := utils.ParsePGIDs(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON("invalid id param")
	}

	orderid, err := utils.ParsePGIDs(deleteReq.OrderId)
	if err != nil {
		return c.Status(400).JSON("invalid order id")
	}

	status := deleteReq.Status
	if status == "" {
		status = "Cancelled"
	}

	params := db.UpdateOrderParams{
		ID:     orderid,
		Status: utils.ToPgTex(&status),
	}

	if err := h.DeleteOrderItemsTx(c.Context(), id, orderid, params); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON("failed to delete order items")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "items successfully deleted",
	})
}
