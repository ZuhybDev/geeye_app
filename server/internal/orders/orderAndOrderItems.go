package orders

import (
	"context"
	"fmt"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/jackc/pgx/v5"
)

func (h *OrderHandler) CreateOrderTx(ctx context.Context, db *pgx.Conn, orderParams db.CreateOrderItemParams, items []db.CreateOrderItemParams) error {
	// 1. Start the transaction using pgxpool
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	// 2. We use a "defer" to ensure we cleanup.
	// If the function returns with an error, Rollback is called.
	// If Commit was already called, Rollback does nothing.
	defer tx.Rollback(ctx)

	// 3. Bind your sqlc queries to this transaction
	qtx := s.queries.WithTx(tx)

	// 4. Create the Order
	order, err := qtx.CreateOrder(ctx, orderParams)
	if err != nil {
		return fmt.Errorf("order creation failed: %w", err)
	}

	// 5. Loop through items
	for _, item := range items {
		item.OrderID = order.ID
		if err := qtx.CreateOrderItem(ctx, item); err != nil {
			// If the 9th item fails, we return here.
			// The 'defer' above will immediately Rollback the whole thing.
			return fmt.Errorf("item creation failed: %w", err)
		}
	}

	// 6. Only if EVERYTHING above succeeded, we commit.
	return tx.Commit(ctx)
}
