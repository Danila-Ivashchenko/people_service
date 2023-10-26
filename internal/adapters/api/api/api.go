package api

import (
	"github.com/gin-gonic/gin"
)

type config interface {
	GetHTTPPort() string
}

type personRouter interface {
	AddPerson(c *gin.Context)
	GetPerson(c *gin.Context)
	UpdatePerson(c *gin.Context)
}

type api struct {
	personRouter personRouter
	server       *gin.Engine
	port         string
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
	a.server.GET("/person", a.personRouter.GetPerson)
	a.server.PATCH("/person", a.personRouter.UpdatePerson)
}

func (a *api) Run() error {
	return a.server.Run(":" + a.port)
}
