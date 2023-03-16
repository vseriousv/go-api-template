package dto

import "github.com/vseriousv/go-api-template/internal/api/repositories"

type DepositDTO struct {
	PoolId      string      `json:"poolId"`
	Total       string      `json:"total"`
	Transaction interface{} `json:"transaction"`
}

func NewDepositDTO(repository repositories.Deposit) *DepositDTO {
	return &DepositDTO{
		PoolId:      repository.PoolId,
		Total:       repository.Total,
		Transaction: "",
	}
}
