package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-achitecture/app/inout"
	errorStruct "github.com/kanok-p/go-clean-achitecture/domain/error"
	serviceUsr "github.com/kanok-p/go-clean-achitecture/service/users"
)

func (app *App) UpdateUsers(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorStruct.ResponseError{
			Error: strconv.Itoa(http.StatusBadRequest),
			Msg:   err.Error(),
		})
		return
	}

	input := inout.User{}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errorStruct.ResponseError{
			Error: strconv.Itoa(http.StatusBadRequest),
			Msg:   err.Error(),
		})
		return
	}

	users := &serviceUsr.UpdateUsers{}
	if err := copier.Copy(users, &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorStruct.ResponseError{
			Error: strconv.Itoa(http.StatusInternalServerError),
			Msg:   err.Error(),
		})
		return
	}
	users.ID = &id

	usersResp, err := app.usrService.Update(ctx, users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorStruct.ResponseError{
			Error: strconv.Itoa(http.StatusInternalServerError),
			Msg:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, usersResp)
}
