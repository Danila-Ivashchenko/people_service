package response

import "people_service/internal/domain/model"

type BadResponse struct {
	Error string `json:"error"`
}

func NewBadResponse(err error) *BadResponse {
	return &BadResponse{
		Error: err.Error(),
	}
}

type IdResponse struct {
	Id int64 `json:"id"`
}

type PersonResponse struct {
	Person model.Person `json:"person"`
}

type PersonsPesponse struct {
	Persons []model.Person `json:"persons"`
}

type SuccessResponse struct {
	Ok bool `json:"ok"`
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{Ok: true}
}
