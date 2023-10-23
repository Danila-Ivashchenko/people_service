package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type configer interface {
	GetPsqlURL() string
}

type postgresClient struct {
	url string
}

func NewPostgresClient(cfg configer) *postgresClient {
	return &postgresClient{
		url: cfg.GetPsqlURL(),
	}
}

func (cl *postgresClient) GetDb() (*sqlx.DB, error) {
	return sqlx.Open("postgres", cl.url)
}
