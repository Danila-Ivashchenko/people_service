package main

import (
	"flag"
	"people_service/pkg/config"
	"people_service/pkg/logger"
	"people_service/pkg/migrator"
)

const (
	actionUp   = "up"
	actionDown = "down"
	actionInit = "init"
)

func main() {

	action := flag.String("action", "", "up or down or init")
	flag.Parse()
	err := config.LoadEnv()
	if err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	logger := logger.SetupLogger(cfg.GetEnv())
	m := migrator.New(cfg, logger)

	switch *action {
	case actionUp:
		err = m.Up()
		if err != nil {
			panic(err)
		}

	case actionDown:
		err = m.Down()
		if err != nil {
			panic(err)
		}
	case actionInit:
		err = m.Init()
        if err!= nil {
            panic(err)
        }
	default:
		panic("invalid action")
	}
}
