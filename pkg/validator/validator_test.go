package validator

import (
	"people_service/internal/domain/dto"
	domain_err "people_service/internal/domain/errors"
	"errors"
	"testing"
)

func TestValidateDataToAdd(t *testing.T) {
	tests := []struct {
		name string
		data *dto.AddPersonRawDTO
		out  error
	}{
		{
			name: "valid",
			data: &dto.AddPersonRawDTO{
				Name:       "Danila",
				Surname:    "Ivashenko",
				Patronymic: "Maksimovich",
			},
			out: nil,
		},
		{
			name: "valid",
			data: &dto.AddPersonRawDTO{
				Name:       "Danila",
				Surname:    "Ivashenko",
				Patronymic: "",
			},
			out: nil,
		},
		{
			name: "invalid name 1",
			data: &dto.AddPersonRawDTO{
				Name:       "danila",
				Surname:    "Ivashenko",
				Patronymic: "",
			},
			out: domain_err.ErrInvalidName,
		},
		{
			name: "invalid name 2",
			data: &dto.AddPersonRawDTO{
				Name:       "DaNila",
				Surname:    "Ivashenko",
				Patronymic: "",
			},
			out: domain_err.ErrInvalidName,
		},
		{
			name: "invalid surname 1",
			data: &dto.AddPersonRawDTO{
				Name:       "Danila",
				Surname:    "ivashenko",
				Patronymic: "",
			},
			out: domain_err.ErrInvalidSurname,
		},
		{
			name: "invalid surname 2",
			data: &dto.AddPersonRawDTO{
				Name:       "Danila",
				Surname:    "IVashenko",
				Patronymic: "",
			},
			out: domain_err.ErrInvalidSurname,
		},
		{
			name: "invalid patronymic 1",
			data: &dto.AddPersonRawDTO{
				Name:       "Danila",
				Surname:    "Ivashenko",
				Patronymic: "Максимович",
			},
			out: domain_err.ErrInvalidPatronymic,
		},
		{
			name: "invalid patronimic 2",
			data: &dto.AddPersonRawDTO{
				Name:       "Danila",
				Surname:    "Ivashenko",
				Patronymic: "MaKsimovich",
			},
			out: domain_err.ErrInvalidPatronymic,
		},
	}

	validator := NewValidator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateDataToAdd(tt.data)
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
		})
	}
}

func TestValidateDataToGet(t *testing.T) {
	tests := []struct{
		name string
		inp *dto.PersonsGetDTO
		out error
	}{
		{
			name: "valid 1",
			inp: &dto.PersonsGetDTO{
				Name: "Oleg",
				Surname: "",
				Patronymic: "",
				Age: 22,
				Gender: "male",
				Nationality: "RU",
			},
			out: nil,
		},
		{
			name: "valid 2",
			inp: &dto.PersonsGetDTO{
				Surname: "Ivashenko",
				Patronymic: "Maksimovich",
				Nationality: "UA",
			},
			out: nil,
		},
		{
			name: "invalid name",
			inp: &dto.PersonsGetDTO{
				Name: "Олeg",
			},
			out: domain_err.ErrInvalidName,
		},
		{
			name: "invalid surname",
			inp: &dto.PersonsGetDTO{
				Surname: "New2",
			},
			out: domain_err.ErrInvalidSurname,
		},
		{
			name: "invalid patronymic",
			inp: &dto.PersonsGetDTO{
				Patronymic: "NewPatr",
			},
			out: domain_err.ErrInvalidPatronymic,
		},
		{
			name: "invalid age",
			inp: &dto.PersonsGetDTO{
				Age: 200,
			},
			out: domain_err.ErrInvalidAge,
		},
		{
			name: "invalid gender",
			inp: &dto.PersonsGetDTO{
				Gender: "Male",
				Limit: 10,
			},
			out: domain_err.ErrInvalidGender,
		},
		{
			name: "invalid nationality",
			inp: &dto.PersonsGetDTO{
				Nationality: "kz",
				Offset: 5,
			},
			out: domain_err.ErrInvalidNationality,
		},
		{
			name: "nothing to get",
			inp: &dto.PersonsGetDTO{
			},
			out: domain_err.ErrNoFiltersToGet,
		},
		{
			name: "invalid limit",
			inp: &dto.PersonsGetDTO{
				Limit: -10,
			},
			out: domain_err.ErrInvalidLimit,
		},
		{
			name: "invalid offset",
			inp: &dto.PersonsGetDTO{
				Offset: -5,
			},
			out: domain_err.ErrInvalidOffset,
		},
	}

	validator := NewValidator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateDataToGet(tt.inp)
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
		})
	}
}


func TestValidateDataToUpdate(t *testing.T) {
	tests := []struct{
		name string
		inp *dto.UpdatePersonDTO
		out error
	}{
		{
			name: "valid 1",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Name: "Oleg",
				Surname: "",
				Patronymic: "",
				Age: 22,
				Gender: "male",
				Nationality: "RU",
			},
			out: nil,
		},
		{
			name: "valid 2",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Surname: "Ivashenko",
				Patronymic: "Maksimovich",
				Nationality: "UA",
			},
			out: nil,
		},
		{
			name: "invalid name",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Name: "Олeg",
			},
			out: domain_err.ErrInvalidName,
		},
		{
			name: "invalid surname",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Surname: "New2",
			},
			out: domain_err.ErrInvalidSurname,
		},
		{
			name: "invalid patronymic",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Patronymic: "NewPatr",
			},
			out: domain_err.ErrInvalidPatronymic,
		},
		{
			name: "invalid age",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Age: 200,
			},
			out: domain_err.ErrInvalidAge,
		},
		{
			name: "invalid gender",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Gender: "Male",
			},
			out: domain_err.ErrInvalidGender,
		},
		{
			name: "invalid nationality",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
				Nationality: "kz",
			},
			out: domain_err.ErrInvalidNationality,
		},
		{
			name: "nothing to update",
			inp: &dto.UpdatePersonDTO{
				Id: 1,
			},
			out: domain_err.ErrNothingToUpdate,
		},
		{
			name: "invalid id 1",
			inp: &dto.UpdatePersonDTO{
				Id: -1,
			},
			out: domain_err.ErrInvalidId,
		},
		{
			name: "invalid id 2",
			inp: &dto.UpdatePersonDTO{
				Id: 0,
			},
			out: domain_err.ErrInvalidId,
		},
	}

	validator := NewValidator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateDataToUpdate(tt.inp)
			if !errors.Is(err, tt.out) {
				t.Errorf("got %v, want %v", err, tt.out)
			}
		})
	}
}
