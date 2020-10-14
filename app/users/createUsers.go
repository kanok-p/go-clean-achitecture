package users

import (
	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (ctrl *Controller) CreateUsers(ctx *gin.Context) {
	input := &inout.Create{}
	if err := ctx.ShouldBind(input); err != nil {
		response.Error(ctx, response.BadRequest(err))
		return
	}

	err := ctrl.service.Create(ctx, input)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Created(ctx, nil)
}
