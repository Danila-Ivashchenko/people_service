package storage

import (
	"context"
	"people_service/internal/domain/dto"
	"people_service/internal/domain/model"
)

type PersonStorage interface {
	AddPerson(context.Context, *dto.AddPersonDTO) (int64, error)
	GetPerson(context.Context, int64) (*model.Person, error)
	UpdatePerson(context.Context, *dto.UpdatePersonDTO) error
	GetPersons(context.Context, *dto.PersonsGetDTO) ([]model.Person, error)
	DeletePerson(context.Context, int64) error
}
