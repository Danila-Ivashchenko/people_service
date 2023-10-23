package service

import (
	"context"
	"fmt"
	"log/slog"
	"people_service/internal/domain/dto"
	"people_service/internal/domain/model"
	"people_service/internal/domain/ports/enricher"
	"people_service/internal/domain/ports/storage"
)

type Validator interface {
	ValidateId(id int64) error
	ValidateDataToAdd(data *dto.AddPersonRawDTO) error
	ValidateDataToGet(data *dto.PersonsGetDTO) error
	ValidateDataToUpdate(data *dto.UpdatePersonDTO) error
}

type Service struct {
	storage  storage.PersonStorage
	enricher enricher.Enricher
	logger   *slog.Logger
	validator Validator
}

func New(s storage.PersonStorage, e enricher.Enricher, l *slog.Logger, v Validator) *Service {
	return &Service{
		storage:  s,
		enricher: e,
		logger:   l,
		validator: v,
	}
}

func (s Service) AddPerson(ctx context.Context, data *dto.AddPersonRawDTO) (int64, error) {
	fmt.Println("here")
	if err := s.validator.ValidateDataToAdd(data); err != nil {
		return -1, err
	}
	fmt.Println("here")
	enricherData, err := s.enricher.Enriche(ctx, data.Name)
	if err != nil {
		return -1, err
	}

	personDto := &dto.AddPersonDTO{
		Name: data.Name,
        Surname: data.Surname,
        Patronymic: data.Patronymic,
        Age: enricherData.Age,
        Gender: enricherData.Gender,
        Nationality: enricherData.Nationality,
	}

	id, err := s.storage.AddPerson(ctx, personDto)
	if err != nil {
		s.logger.Debug("Failed to add person", slog.Any("error", err))
		return -1, err
	}
	fmt.Println("here")
	s.logger.Debug("Person added", slog.Any("id", id))
	return id, nil
}

func (s Service) GetPerson(ctx context.Context, id int64) (*model.Person, error) {
	if err := s.validator.ValidateId(id); err != nil {
		return nil, err
	}
	return s.storage.GetPerson(ctx, id)
}

func (s Service) DeletePerson(ctx context.Context, id int64) error {
	if err := s.validator.ValidateId(id); err!= nil {
        return err
    }
    return s.storage.DeletePerson(ctx, id)
}

func (s Service) UpdatePerson(ctx context.Context, data *dto.UpdatePersonDTO) error {
	if err := s.validator.ValidateDataToUpdate(data); err != nil {
		return err
	}

	return s.storage.UpdatePerson(ctx, data)
}