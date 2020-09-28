package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-architecture/domain/request"
	"github.com/kanok-p/go-clean-architecture/domain/response"
)

func (app *App) ListUsers(ctx *gin.Context) {
	input := &request.GetListInput{}
	if err := ctx.ShouldBind(input); err != nil {
		response.Error(ctx, response.BadRequest(err))
		return
	}

	input.Limit = input.GetLimit()
	if input.Offset == 0 {
		input.Offset = input.GetOffset()
	}

	total, list, err := app.usrService.List(ctx, input)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &response.ListResp{
		Pagination: response.Pagination{
			Total:  total,
			Page:   input.GetPage(),
			Limit:  input.Limit,
			Search: input.Search,
			Sort:   "",
		},
		List: list,
	})
}
