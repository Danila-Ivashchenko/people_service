package storage

import (
	"context"
	"people_service/internal/domain/dto"
	domain_err "people_service/internal/domain/errors"
	"people_service/internal/domain/model"

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
	stmt += " RETURNING id"
	var id int64
	insertStmt, err := tx.PrepareNamedContext(ctx, stmt)
	if err != nil {
		return -1, err
	}

	err = insertStmt.GetContext(ctx, &id, data)
	if err != nil {
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, err
	}

	return id, nil
}

func (s *personStorage) GetPerson(ctx context.Context, id int64) (*model.Person, error) {
	stmt := "SELECT id, name, surname, patronymic, age, gender, nationality FROM persons WHERE id = $1"
	person := &model.Person{}
	err := s.db.QueryRowxContext(ctx, stmt, id).StructScan(person)
	if err != nil {
		return nil, domain_err.ErrorNoSuchUser
	}

	return person, nil
}
func (s *personStorage) UpdatePerson(ctx context.Context, data *dto.UpdatePersonDTO) error {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := "UPDATE persons SET " + data.ExtractSQL() + " WHERE id = :id"
	updateStmt, err := tx.PrepareNamedContext(ctx, stmt)
	if err != nil {
		return err
	}

	result, err := updateStmt.Exec(data)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if  rows == 0 {
		return domain_err.ErrorNoSuchUser
	}

	return tx.Commit()
}
func (s *personStorage) GetPersons(context.Context, *dto.PersonsGetDTO) ([]model.Person, error) {
	return nil, nil
}
func (s *personStorage) DeletePerson(context.Context, int64) error {
	return nil
}
