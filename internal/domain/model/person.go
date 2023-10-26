package model

const (
	MALE   = "male"
	FEMALE = "female"
)

type Person struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         uint   `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}
