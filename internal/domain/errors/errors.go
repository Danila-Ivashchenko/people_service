package errors

import "errors"

var (
	ErrInvalidId          = errors.New("invalid id")
	ErrNoProvidedId        = errors.New("id was not provided, but required")

	ErrInvalidName        = errors.New("invalid name")
	ErrNoProvidedName        = errors.New("name was not provided, but required")

	ErrInvalidSurname     = errors.New("invalid surname")
	ErrNoProvidedSurname = errors.New("surname was not provided, but required")

	ErrInvalidPatronymic  = errors.New("invalid patronymic")
	ErrNoProvidedPatronymic = errors.New("patronymic was not provided, but required")

	ErrInvalidAge         = errors.New("invalid age")
	ErrInvalidGender      = errors.New("invalid gender")
	ErrInvalidNationality = errors.New("invalid nationality")

	ErrDatabse  = errors.New("databse error")
	ErrEnricher = errors.New("enricher error")
)
