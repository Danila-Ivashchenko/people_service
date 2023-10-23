package storage

import "github.com/jmoiron/sqlx"

type psqlClinet interface {
	GetDb() (*sqlx.DB, error)
}

type personStorage struct {
	psqlClinet psqlClinet
}

func NewPersonStorage(p psqlClinet) *personStorage {
	return &personStorage{
		psqlClinet: p,
	}
}
