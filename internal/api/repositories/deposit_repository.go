package repositories

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Deposit struct {
	PoolId string `json:"poolId"`
	Total  string `json:"total"`
}

type DepositRepository struct {
	db *pgxpool.Pool
}

func NewDepositRepository(pool *pgxpool.Pool) *DepositRepository {
	return &DepositRepository{db: pool}
}

func (r *DepositRepository) CalculateInvest(ctx context.Context, d *Deposit) error {
	return nil
}
