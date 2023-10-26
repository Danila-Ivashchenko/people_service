package main

import (
	"people_service/internal/adapters/enricher"
	"people_service/internal/adapters/storage"
	"people_service/internal/domain/service"
	"people_service/pkg/config"
	"people_service/pkg/logger"
	"people_service/pkg/psql"
	"people_service/pkg/validator"
	"people_service/internal/adapters/api/router"
	"people_service/internal/adapters/api/api"

)

func main() {
	if err := config.LoadEnv(); err != nil {
		panic(err)
	}

	cfg := config.GetConfig()
	logger := logger.SetupLogger(cfg.GetEnv())

	enricher := enricher.New(cfg)
	psqlClient := psql.NewPostgresClient(cfg)
	db, err := psqlClient.GetDb()
	if err != nil {
		panic(err)
	}
	personStorage := storage.NewPersonStorage(db)

	personService := service.New(personStorage, enricher, logger, validator.NewValidator())
	personRouter := router.NewPersonRouter(personService)
	app := api.New(cfg, personRouter)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
