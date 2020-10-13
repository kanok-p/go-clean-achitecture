package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/app/users"
)

type App struct {
	controller users.Controller
}

func New(usrController users.Controller) *App {
	return &App{
		controller: usrController,
	}
}

func (app *App) Register(router *gin.Engine) *App {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	router.GET("/users", app.controller.ListUsers)
	router.POST("/users", app.controller.CreateUsers)
	router.GET("/users/:id", app.controller.GetUsers)
	router.PUT("/users/:id", app.controller.UpdateUsers)
	router.DELETE("/users/:id", app.controller.DeleteUsers)

	return app
}
