package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	errorStruct "github.com/kanok-p/go-clean-achitecture/domain/error"
)

func (app *App) GetUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := app.usrService.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorStruct.ResponseError{
			Error: strconv.Itoa(http.StatusNotFound),
			Msg:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, users)
}
