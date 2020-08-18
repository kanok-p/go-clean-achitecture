package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	errorStruct "github.com/kanok-p/go-clean-achitecture/domain/error"
)

func (app *App) DeleteUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := app.usrService.Delete(ctx, id)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			ctx.JSON(http.StatusBadRequest, errorStruct.ResponseError{
				Error: strconv.Itoa(http.StatusBadRequest),
				Msg:   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusNotFound, errorStruct.ResponseError{
				Error: strconv.Itoa(http.StatusNotFound),
				Msg:   err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, users)
}
