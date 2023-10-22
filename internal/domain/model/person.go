package model

const (
	MALE = "male"
	FEMALE = "female"
)

type Person struct {
	Id int64
	Name string
	Surname string
	Patronymic string
	Age uint
	Gender string
	Nationality string
}

