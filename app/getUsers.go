package app

import (
	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/domain/response"
)

func (app *App) GetUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := app.usrService.Get(ctx, id)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.OK(ctx, users)
}
