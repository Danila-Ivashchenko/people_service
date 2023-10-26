package service

import (
	"context"
	"people_service/internal/domain/dto"
	"people_service/internal/domain/model"
)

type PersonService interface {
	AddPerson(ctx context.Context, data *dto.AddPersonRawDTO) (int64, error)
	GetPerson(ctx context.Context, id int64) (*model.Person, error)
	GetPersons(ctx context.Context, data *dto.PersonsGetDTO) ([]model.Person, error)
	UpdatePerson(ctx context.Context, data *dto.UpdatePersonDTO) error
	DeletePerson(ctx context.Context, id int64) error
}
