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
	url string
	logger *slog.Logger
}

func New(cfg configer, l *slog.Logger) *migrater {
	return &migrater{
		url: cfg.GetPsqlURL(),
		logger: l,
	}
}

func (mig migrater) Migrate() error {
	op := "migrator/migrate"
	logger := mig.logger.With("operation", op)

	m, err := migrate.New("file://migrations", mig.url)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Error("fail to migrate", slog.String("error", err.Error()))
        return err
    }
	v, d, err := m.Version()
    if err!= nil {
		logger.Error("fail to get version", slog.String("error", err.Error()))
        return err
    }
	logger.Info("successfully to migrate", slog.String("version", fmt.Sprintf("%d", v)), slog.String("is dirty", fmt.Sprintf("%t", d)))

	return nil
}
