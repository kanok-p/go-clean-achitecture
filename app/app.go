package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/app/users"
)

type App struct {
	usrController users.Controller
}

func New(usrController users.Controller) *App {
	return &App{
		usrController: usrController,
	}
}

func (app *App) Register(router *gin.Engine) *App {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	router.GET("/users", app.usrController.ListUsers)
	router.POST("/users", app.usrController.CreateUsers)
	router.GET("/users/:id", app.usrController.GetUsers)
	router.PUT("/users/:id", app.usrController.UpdateUsers)
	router.DELETE("/users/:id", app.usrController.DeleteUsers)

	return app
}
