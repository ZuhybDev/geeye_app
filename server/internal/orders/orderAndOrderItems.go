package orders

import (
	"context"
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *OrderHandler) CreateOrderTx(ctx context.Context, orderParams db.CreateOrderParams, items []db.CreateOrderItemParams) error {
	// 1. Start the transaction using pgxpool

	tx, err := h.DbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("DB connection failed: %w", err)
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

func (h *OrderHandler) DeleteOrderItemsTx(ctx context.Context, id pgtype.UUID, orderId pgtype.UUID, updateOrder db.UpdateOrderParams) error {

	tx, err := h.DbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	qtx := h.Cfg.Query.WithTx(tx)

	if err := qtx.DeleteOrderItems(ctx, id); err != nil {
		return fmt.Errorf("delete order items: %w", err)
	}

	hasItems, err := qtx.CheckOrderItems(ctx, orderId)
	if err != nil {
		return fmt.Errorf("check order items: %w", err)
	}

	if !hasItems {
		if _, err := qtx.UpdateOrder(ctx, updateOrder); err != nil {
			return fmt.Errorf("update order: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
