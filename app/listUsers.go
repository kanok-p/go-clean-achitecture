package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
	"github.com/kanok-p/go-clean-achitecture/util/pagination"
)

type GetListInput struct {
	*pagination.Pagination
	Limit  int64
	Offset int64
}

type ListResp struct {
	Total int64
	List  []*domainUsr.Users
}

func (app *App) ListUsers(ctx *gin.Context) {

	input := &GetListInput{}
	if err := ctx.ShouldBind(input); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	total, list, err := app.usrService.List(ctx, input.Offset, input.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, &ListResp{
		Total: total,
		List:  list,
	})
}
