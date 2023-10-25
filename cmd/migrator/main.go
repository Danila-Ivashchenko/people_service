package main

import (
	"people_service/pkg/config"
	"people_service/pkg/logger"
	"people_service/pkg/migrator"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	logger := logger.SetupLogger(cfg.GetEnv())
	m := migrator.New(cfg, logger)
	err = m.Migrate()
	if err != nil {
		panic(err)
	}
}
