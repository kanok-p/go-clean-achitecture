package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/domain/request"
	"github.com/kanok-p/go-clean-architecture/domain/response"
)

func (ctrl *Controller) ListUsers(ctx *gin.Context) {
	input := &request.PageOption{}
	if err := ctx.ShouldBind(input); err != nil {
		response.Error(ctx, response.BadRequest(err))
		return
	}

	total, list, err := ctrl.service.List(ctx, input)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &response.ListResp{
		Pagination: response.Pagination{
			Total:   total,
			Page:    input.Page,
			PerPage: input.PerPage,
			Search:  input.Search,
		},
		List: list,
	})
}
