package users

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (ctrl *Controller) UpdateUsers(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		response.Error(ctx, response.BadRequest(err))
		return
	}

	input := &inout.Update{}
	if err := ctx.ShouldBind(&input); err != nil {
		response.Error(ctx, response.BadRequest(err))
		return
	}

	input.ID = &id
	user, err := ctrl.service.Update(ctx, input)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.OK(ctx, user)
}
