package router

import (
	"context"
	"fmt"
	"net/http"
	"people_service/internal/adapters/api/response"
	"people_service/internal/domain/dto"
	domain_err "people_service/internal/domain/errors"
	"people_service/internal/domain/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type service interface {
	AddPerson(ctx context.Context, data *dto.AddPersonRawDTO) (int64, error)
	GetPerson(ctx context.Context, id int64) (*model.Person, error)
	UpdatePerson(ctx context.Context, data *dto.UpdatePersonDTO) error
}

type personRouter struct {
	service service
}

func NewPersonRouter(s service) *personRouter {
	return &personRouter{
		service: s,
	}
}

func (r personRouter) AddPerson(c *gin.Context) {
	request := &dto.AddPersonRawDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	id, err := r.service.AddPerson(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}
	c.JSON(http.StatusOK, response.IdResponse{Id: id})
}

func (r personRouter) GetPerson(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(fmt.Errorf("no id in params")))
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(domain_err.ErrInvalidId))
		return
	}

	person, err := r.service.GetPerson(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}
	c.JSON(http.StatusOK, person)
}

func (r personRouter) UpdatePerson(c *gin.Context) {
	request := &dto.UpdatePersonDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	err := r.service.UpdatePerson(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"ok": true})
}
