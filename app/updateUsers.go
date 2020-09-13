package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/app/inout"
	"github.com/kanok-p/go-clean-architecture/domain/response"
	serviceUsr "github.com/kanok-p/go-clean-architecture/service/users"
)

func (app *App) UpdateUsers(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, err)
		return
	}

	input := inout.User{}
	if err := ctx.ShouldBind(&input); err != nil {
		response.Error(ctx, err)
		return
	}

	users := &serviceUsr.UpdateUsers{}
	if err := copier.Copy(users, &input); err != nil {
		response.Error(ctx, err)
		return
	}
	users.ID = &id

	usersResp, err := app.usrService.Update(ctx, users)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.OK(ctx, usersResp)
}
