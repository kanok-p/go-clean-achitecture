package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kanok-p/go-clean-achitecture/domain/request"
	"github.com/kanok-p/go-clean-achitecture/domain/response"
)

func (app *App) ListUsers(ctx *gin.Context) {
	input := &request.GetListInput{}
	if err := ctx.ShouldBind(input); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	input.Limit = input.GetLimit()
	if input.Offset == 0 {
		input.Offset = input.GetOffset()
	}

	total, list, err := app.usrService.List(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
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
