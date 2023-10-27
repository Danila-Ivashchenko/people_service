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

// AddPerson
// @Summary add person to data base
// @Tags person
// @Description get NSP to enriche it and add
// @ID add-person
// @Accept json
// @Produce json
// @Param input body dto.AddPersonRawDTO true "name, surname, patronymic"
// @Success 201 {object} response.IdResponse
// @Failure 400 {object} response.BadResponse
// @Router /person [post]
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
	c.JSON(http.StatusCreated, response.IdResponse{Id: id})
}

// GetPerson
// @Summary get person by id in fromto data base
// @Tags person
// @Description get id from url params and find person
// @ID get-person
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} model.Person
// @Failure 400 {object} response.BadResponse
// @Router /person/{id} [get]
func (r personRouter) GetPerson(c *gin.Context) {
	idStr := c.Param("id")
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

// GetPersons
// @Summary get a list of persons by params and pagination
// @Tags person
// @Description Get a list of persons based on query parameters
// @ID get-persons
// @Accept  json
// @Produce  json
// @Param name query string false "name of the person"
// @Param surname query string false "surname of the person"
// @Param patronymic query string false "patronymic of the person"
// @Param age query int false "age of the person"
// @Param gender query string false "gender of the person"
// @Param nationality query string false "nationality of the person"
// @Param limit query int false "limit the number of results"
// @Param offset query int false "offset for pagination"
// @Success 200 {object} []model.Person
// @Failure 400 {object} response.BadResponse
// @Router /persons [get]
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

// UpdatePerson
// @Summary update person in data base
// @Tags person
// @Description update person
// @ID update-person
// @Accept json
// @Produce json
// @Param input body dto.UpdatePersonDTO true "id, name, surname, patronymic, age, gender, nationality"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.BadResponse
// @Router /person [patch]
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
	c.JSON(http.StatusOK, response.NewSuccessResponse())
}

// DeletePerson
// @Summary delete person from data base
// @Tags person
// @Description delete person
// @ID delete-person
// @Accept json
// @Produce json
// @Param input body dto.IdDTO true "id"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.BadResponse
// @Router /person [delete]
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
	c.JSON(http.StatusOK, response.NewSuccessResponse())
}
