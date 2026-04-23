package orders

import (
	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderHandler struct {
	Cfg    *env.Config
	DbPool *pgxpool.Pool
}

func NewOrderHandler(a *env.Config, myPool *pgxpool.Pool) *OrderHandler {
	return &OrderHandler{
		Cfg:    a,
		DbPool: myPool,
	}
}
