package services

import (
	"backend/internal/api/dto"
	"backend/internal/api/repositories"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DepositService struct {
	db *pgxpool.Pool
}

func NewDepositService(db *pgxpool.Pool) *DepositService {
	return &DepositService{
		db: db,
	}
}

func (r *DepositService) GetDeposit() (*dto.DepositDTO, error) {
	repository := repositories.NewDepositRepository(r.db)
	depositEntity := repositories.Deposit{}
	err := repository.CalculateInvest(context.Background(), &depositEntity)
	if err != nil {
		return nil, err
	}

	deposit := dto.NewDepositDTO(depositEntity)
	return deposit, nil
}
