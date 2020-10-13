package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	serviceUsr "github.com/kanok-p/go-clean-architecture/service/users"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (ctrl *Controller) CreateUsers(ctx *gin.Context) {
	input := inout.User{}
	if err := ctx.ShouldBind(&input); err != nil {
		response.Error(ctx, response.BadRequest(err))
		return
	}

	users := &serviceUsr.CreateUsers{}
	if err := copier.Copy(users, &input); err != nil {
		response.Error(ctx, response.InternalServerError(err))
		return
	}

	err := ctrl.service.Create(ctx, users)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Created(ctx, users)
}
