package users

import (
	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/domain/response"
)

func (ctrl *Controller) DeleteUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := ctrl.service.Delete(ctx, id)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.OK(ctx, users)
}
