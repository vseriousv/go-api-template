package api

import (
	"backend/internal/api/handlers"
	"backend/internal/api/services"
	"backend/internal/config"
	"backend/pkg/utils"
	"context"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AppStruct struct {
	Router chi.Router
}

func (a *AppStruct) RunApp(c *config.Config) {
	configPgx, err := pgxpool.ParseConfig(c.DbUrl)
	utils.HandleError(err, "DB is not connection")
	configPgx.MinConns = 64
	configPgx.MaxConns = 256
	dbContext := context.Background()
	db, err := pgxpool.ConnectConfig(dbContext, configPgx)
	utils.HandleError(err, "DB is not connection")
	err = db.Ping(context.Background())
	utils.HandleError(err, "DB is not connection")

	router := chi.NewRouter()

	// services with db
	depositService := services.NewDepositService(db)

	// handlers
	depositHandler := handlers.NewDepositHandler(depositService)

	// GET DEPOSIT CALCULATE
	router.Get("/v1/deposit", depositHandler.DepositCalculate)
}
