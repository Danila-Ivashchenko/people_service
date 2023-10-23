package validator

import (
	"fmt"
	"people_service/internal/domain/dto"
	domain_err "people_service/internal/domain/errors"
	"people_service/internal/domain/model"
	"regexp"

	"github.com/pkg/errors"
)

type validator struct {
	re            *regexp.Regexp
	reNationality *regexp.Regexp
}

func NewValidator() *validator {
	return &validator{
		re:            regexp.MustCompile(`^[A-Z][a-z]+$`),
		reNationality: regexp.MustCompile(`^[A-Z]{2}$`),
	}
}

func (validator) ValidateId(id int64) error {
	if id <= 0 {
		return errors.Wrap(domain_err.ErrInvalidId, fmt.Sprintf("%d", id))
	}
	return nil
}

func (v validator) ValidateName(name string) error {
	if v.re.MatchString(name) {
		return nil
	} else {
		return errors.Wrap(domain_err.ErrInvalidName, name)
	}
}
func (v validator) ValidateSurame(surname string) error {
	if v.re.MatchString(surname) {
		return nil
	} else {
		return errors.Wrap(domain_err.ErrInvalidSurname, surname)
	}
}
func (v validator) ValidatePatronymic(patronymic string) error {
	if v.re.MatchString(patronymic) {
		return nil
	} else {
		return errors.Wrap(domain_err.ErrInvalidPatronymic, patronymic)
	}
}

func (v validator) ValidateDataToAdd(data *dto.AddPersonRawDTO) error {
	if err := v.ValidateName(data.Name); err != nil {
		return err
	}
	if err := v.ValidateSurame(data.Surname); err != nil {
		return err
	}
	if data.Patronymic != "" {
		if err := v.ValidatePatronymic(data.Patronymic); err != nil {
			return err
		}
	}

	return nil
}

func (v validator) ValidateGender(gender string) error {
	if gender == model.MALE || gender == model.FEMALE {
		return nil
	} else {
		return errors.Wrap(domain_err.ErrInvalidGender, gender)
	}
}

func (v validator) ValidateNationality(nationality string) error {
	if v.reNationality.MatchString(nationality) {
		return nil
	} else {
		return errors.Wrap(domain_err.ErrInvalidNationality, nationality)
	}
}

func (v validator) ValidateAge(age uint) error {
	if age > 150 {
		return errors.Wrap(domain_err.ErrInvalidAge, fmt.Sprintf("%d", age))
	}
	return nil
}

func (v validator) ValidateDataToGet(data *dto.PersonsGetDTO) error {
	fields := 0
	if data.Name != "" {
		if err := v.ValidateName(data.Name); err != nil {
			return err
		}
		fields++
	}

	if data.Surname != "" {
		if err := v.ValidateSurame(data.Surname); err != nil {
			return err
		}
		fields++
	}

	if data.Patronymic != "" {
		if err := v.ValidatePatronymic(data.Patronymic); err != nil {
			return err
		}
		fields++
	}

	if data.Age != 0 {
		if err := v.ValidateAge(data.Age); err != nil {
			return err
		}
		fields++
	}

	if data.Gender != "" {
		if err := v.ValidateGender(data.Gender); err != nil {
			return err
		}
		fields++
	}

	if data.Nationality != "" {
		if err := v.ValidateNationality(data.Nationality); err != nil {
			return err
		}
		fields++
	}

	if data.Limit < 0 {
		return errors.Wrap(domain_err.ErrInvalidLimit, fmt.Sprintf("%d", data.Limit))
	}

	if data.Offset < 0 {
		return errors.Wrap(domain_err.ErrInvalidOffset, fmt.Sprintf("%d", data.Offset))
	}

	if fields == 0 {
		return domain_err.ErrNoFiltersToGet
	}
	return nil

}
func (v validator) ValidateDataToUpdate(data *dto.UpdatePersonDTO) error {
	fields := 0
	if err := v.ValidateId(data.Id); err != nil {
		return err
	}

	if data.Name != "" {
		if err := v.ValidateName(data.Name); err != nil {
			return err
		}
		fields++
	}

	if data.Surname != "" {
		if err := v.ValidateSurame(data.Surname); err != nil {
			return err
		}
		fields++
	}

	if data.Patronymic != "" {
		if err := v.ValidatePatronymic(data.Patronymic); err != nil {
			return err
		}
		fields++
	}

	if data.Age != 0 {
		if err := v.ValidateAge(data.Age); err != nil {
			return err
		}
		fields++
	}

	if data.Gender != "" {
		if err := v.ValidateGender(data.Gender); err != nil {
			return err
		}
		fields++
	}

	if data.Nationality != "" {
		if err := v.ValidateNationality(data.Nationality); err != nil {
			return err
		}
		fields++
	}

	if fields == 0 {
		return domain_err.ErrNothingToUpdate
	}
	return nil
}
