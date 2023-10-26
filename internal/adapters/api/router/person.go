package router

import (
	"fmt"
	"net/http"
	"people_service/internal/adapters/api/response"
	"people_service/internal/adapters/api/service"
	"people_service/internal/domain/dto"
	domain_err "people_service/internal/domain/errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type personRouter struct {
	service service.PersonService
}

func NewPersonRouter(s service.PersonService) *personRouter {
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
	c.JSON(http.StatusCreated, person)
}

func (r personRouter) GetPersons(c *gin.Context) {
	data := &dto.PersonsGetDTO{}
	data.Name = c.Query("name")
	data.Surname = c.Query("surname")
	data.Patronymic = c.Query("patronymic")
	ageStr := c.Query("age")
	if ageStr != "" {
		ageInt, err := strconv.Atoi(ageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(errors.Wrap(domain_err.ErrInvalidAge, ageStr)))
			return
		}
		if ageInt < 0 {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(errors.Wrap(domain_err.ErrInvalidAge, fmt.Sprintf("%d", ageInt))))
			return
		}
		data.Age = uint(ageInt)
	}
	data.Gender = c.Query("gender")
	data.Nationality = c.Query("nationality")
	limitStr := c.Query("limit")
	if limitStr != "" {
		var err error
		data.Limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(errors.Wrap(domain_err.ErrInvalidLimit, limitStr)))
			return
		}
	}
	offsetStr := c.Query("offset")
	if offsetStr != "" {
		var err error
		data.Offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(errors.Wrap(domain_err.ErrInvalidOffset, offsetStr)))
			return
		}
	}

	persons, err := r.service.GetPersons(c.Request.Context(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}
	c.JSON(http.StatusOK, persons)
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

func (r personRouter) DeletePerson(c *gin.Context) {
	request := &dto.IdDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	err := r.service.DeletePerson(c.Request.Context(), request.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}
	c.JSON(http.StatusOK, map[string]bool{"ok": true})
}
