package services

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vseriousv/go-api-template/internal/api/dto"
	"github.com/vseriousv/go-api-template/internal/api/repositories"
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
