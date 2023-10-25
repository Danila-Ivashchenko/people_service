package router

import (
	"context"
	"people_service/internal/domain/dto"
)

type service interface {
	AddPerson(ctx context.Context, data *dto.AddPersonRawDTO) (int64, error)
}

type personRouter struct {
	service service
}

func NewPersonRouter(s service) *personRouter {
	return &personRouter{
		service: s,
	}
}

func (r personRouter) AddPerson(c gin.Context) {

}
