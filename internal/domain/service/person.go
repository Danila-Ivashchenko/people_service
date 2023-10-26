package service

import (
	"context"
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
	storage   storage.PersonStorage
	enricher  enricher.Enricher
	logger    *slog.Logger
	validator Validator
}

func New(s storage.PersonStorage, e enricher.Enricher, l *slog.Logger, v Validator) *Service {
	return &Service{
		storage:   s,
		enricher:  e,
		logger:    l,
		validator: v,
	}
}

func (s Service) AddPerson(ctx context.Context, data *dto.AddPersonRawDTO) (int64, error) {
	op := "service/AddPerson"
	logger := s.logger.With("operation", op)
	if err := s.validator.ValidateDataToAdd(data); err != nil {
		return -1, err
	}

	enricherData, err := s.enricher.Enriche(ctx, data.Name)
	if err != nil {
		return -1, err
	}

	personDto := &dto.AddPersonDTO{
		Name:        data.Name,
		Surname:     data.Surname,
		Patronymic:  data.Patronymic,
		Age:         enricherData.Age,
		Gender:      enricherData.Gender,
		Nationality: enricherData.Nationality,
	}

	id, err := s.storage.AddPerson(ctx, personDto)
	if err != nil {
		logger.Debug("Failed to add person", slog.Any("error", err))
		return -1, err
	}

	logger.Debug("Person added", slog.Any("id", id))
	return id, nil
}

func (s Service) GetPerson(ctx context.Context, id int64) (*model.Person, error) {
	op := "service/GetPerson"
	logger := s.logger.With("operation", op)
	if err := s.validator.ValidateId(id); err != nil {
		return nil, err
	}
	person, err := s.storage.GetPerson(ctx, id)
	if err != nil {
		logger.Error("Faild to get person", slog.Int64("id", id), slog.Any("error", err))
		return nil, err
	}
	logger.Debug("Successly got person", slog.Int64("id", id))
	return person, nil
}

func (s Service) DeletePerson(ctx context.Context, id int64) error {
	if err := s.validator.ValidateId(id); err != nil {
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
