package storage

import (
	"context"
	"people_service/internal/domain/dto"

	"github.com/jmoiron/sqlx"
)

type psqlClinet interface {
	GetDb() (*sqlx.DB, error)
}

type personStorage struct {
	// psqlClinet psqlClinet
	db *sqlx.DB
}

func NewPersonStorage(db *sqlx.DB) *personStorage {
	return &personStorage{
		// psqlClinet: p,
		db: db,
	}
}

func (s *personStorage) AddPerson(ctx context.Context, data *dto.AddPersonDTO) (int64, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	stmt := `INSERT INTO persons (name, surname, patronymic, age, gender, nationality) VALUES (:name, :surname, :patronymic, :age, :gender, :nationality)`
	result, err := tx.NamedExec(stmt, data)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}
	
	return id, nil
}
