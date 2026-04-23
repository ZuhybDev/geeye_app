package orders

import (
	"context"
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
)

func (h *OrderHandler) CreateOrderTx(ctx context.Context, orderParams db.CreateOrderParams, items []db.CreateOrderItemParams) error {
	// 1. Start the transaction using pgxpool

	tx, err := h.DbPool.Begin(ctx)
	if err != nil {
		return err
	}

	// If Commit was already called, Rollback does nothing.
	defer tx.Rollback(ctx)

	qtx := h.Cfg.Query.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, orderParams)

	if err != nil {
		return fmt.Errorf("Order creation failed: %w", err)
	}

	for _, i := range items {
		i.OrderID = order.ID
		if err := qtx.CreateOrderItem(ctx, i); err != nil {
			return fmt.Errorf("Item creation failed: %w", err)
		}
	}

	return tx.Commit(ctx)
}
