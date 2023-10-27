package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "people_service/docs"
)

type config interface {
	GetHTTPPort() string
	GetEnv() string
}

type personRouter interface {
	AddPerson(c *gin.Context)
	GetPerson(c *gin.Context)
	UpdatePerson(c *gin.Context)
	DeletePerson(c *gin.Context)
	GetPersons(c *gin.Context)
}

type api struct {
	personRouter personRouter
	server       *gin.Engine
	port         string
	env          string
}

func New(cfg config, p personRouter) *api {
	api := &api{
		personRouter: p,
		port:         cfg.GetHTTPPort(),
	}

	api.server = gin.New()
	api.server.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/pathsNotToLog/"),
		gin.Recovery(),
	)
	api.bind()

	return api
}

func (a *api) bind() {
	a.server.POST("/person", a.personRouter.AddPerson)
	a.server.GET("/person/:id", a.personRouter.GetPerson)
	a.server.PATCH("/person", a.personRouter.UpdatePerson)
	a.server.DELETE("/person", a.personRouter.DeletePerson)
	a.server.GET("/persons", a.personRouter.GetPersons)

	if a.env != "prod" {
		a.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (a *api) Run() error {
	return a.server.Run(":" + a.port)
}
