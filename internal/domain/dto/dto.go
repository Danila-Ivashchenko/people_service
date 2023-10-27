package dto

type AddPersonRawDTO struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type AddPersonDTO struct {
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         uint   `db:"age"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
}

type UpdatePersonDTO struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         uint   `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
}

func (dto UpdatePersonDTO) ExtractSQL() string {
	result := ""
	if dto.Name != "" {
		result += "name = :name"
	}
	if dto.Surname != "" {
		if result != "" {
			result += ", "
		}
		result += "surname = :surname"
	}
	if dto.Patronymic != "" {
		if result != "" {
			result += ", "
		}
		result += "patronymic = :patronymic"
	}
	if dto.Age != 0 {
		if result != "" {
			result += ", "
		}
		result += "age = :age"
	}
	if dto.Gender != "" {
		if result != "" {
			result += ", "
		}
		result += "gender = :gender"
	}
	if dto.Nationality != "" {
		if result != "" {
			result += ", "
		}
		result += "nationality = :nationality"
	}
	return result
}

type PersonsGetDTO struct {
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         uint   `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
	Limit       int    `json:"limit" db:"limit"`
	Offset      int    `json:"offset" db:"offset"`
}

func (dto PersonsGetDTO) ExtractSQL() string {
	result := ""
	if dto.Name != "" {
		result += "name = :name"
	}
	if dto.Surname != "" {
		if result != "" {
			result += ", "
		}
		result += "surname = :surname"
	}
	if dto.Patronymic != "" {
		if result != "" {
			result += ", "
		}
		result += "patronymic = :patronymic"
	}
	if dto.Age != 0 {
		if result != "" {
			result += ", "
		}
		result += "age = :age"
	}
	if dto.Gender != "" {
		if result != "" {
			result += ", "
		}
		result += "gender = :gender"
	}
	if dto.Nationality != "" {
		if result != "" {
			result += ", "
		}
		result += "nationality = :nationality"
	}

	result += " ORDER BY id"

	if dto.Limit != 0 {
		result += " LIMIT :limit"
	}

	if dto.Offset != 0 {
		result += " OFFSET :offset"
	}

	return result
}

type EnrichDataDTO struct {
	Age         uint
	Gender      string
	Nationality string
}

type AgeDTO struct {
	Age uint `json:"age"`
}

type GenderDTO struct {
	Gender string `json:"gender"`
}

type NationalityDTO struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type NationalitiesDTO struct {
	Country []NationalityDTO `json:"country"`
}

type IdDTO struct {
	Id int64 `json:"id" db:"id"`
}
