package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-achitecture/service/users"
)

type App struct {
	usrService users.Service
}

func New(usrService users.Service) *App {
	return &App{
		usrService: usrService,
	}
}

func (app *App) Register(router *gin.Engine) *App {
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	router.GET("/users", app.ListUsers)
	router.POST("/users", app.CreateUsers)
	router.GET("/users/:id", app.GetUsers)
	router.PUT("/users/:id", app.UpdateUsers)
	router.DELETE("/users/:id", app.DeleteUsers)

	return app
}
