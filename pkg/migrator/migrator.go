package migrator

import (
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type configer interface {
	GetPsqlURL() string
}

type migrater struct {
	url    string
	logger *slog.Logger
}

func New(cfg configer, l *slog.Logger) *migrater {
	return &migrater{
		url:    cfg.GetPsqlURL(),
		logger: l,
	}
}

func (mig migrater) Up() error {
	op := "migrator/up"
	logger := mig.logger.With("operation", op)

	m, err := migrate.New("file://migrations", mig.url)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Error("fail to up", slog.String("error", err.Error()))
		return err
	}
	v, d, err := m.Version()
	if err != nil {
		logger.Error("fail to get version", slog.String("error", err.Error()))
		return err
	}
	logger.Info("successfully to up", slog.String("version", fmt.Sprintf("%d", v)), slog.String("is dirty", fmt.Sprintf("%t", d)))

	return nil
}

func (mig migrater) Down() error {
	op := "migrator/down"
	logger := mig.logger.With("operation", op)

	m, err := migrate.New("file://migrations", mig.url)
	if err != nil {
		return err
	}
	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		logger.Error("fail to down", slog.String("error", err.Error()))
		return err
	}
	// v, d, err := m.Version()
	// if err != nil {
	// 	logger.Error("fail to get version", slog.String("error", err.Error()))
	// 	return err
	// }
	// logger.Info("successfully to down", slog.String("version", fmt.Sprintf("%d", v)), slog.String("is dirty", fmt.Sprintf("%t", d)))
	logger.Info("successfully to down")

	return nil
}
