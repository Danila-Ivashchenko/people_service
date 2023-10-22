package dto

type AddPersonRawDTO struct {
	Name string
	Surname string
	Patronymic string
}

type AddPersonDTO struct {
	Name string
	Surname string
	Patronymic string
	Age uint
	Gender string
	Nationality string
}

type UpdatePersonDTO struct {
	Id int64
	Name string
	Surname string
	Patronymic string
	Age uint
	Gender string
	Nationality string
}

type PersonsGetDTO struct {
	Name string
	Surname string
	Patronymic string
	Age uint
	Gender string
	Nationality string
	Limit int
	Offset int
}

type EnrichDataDTO struct {
	Age uint
	Gender string
	Nationality string
}