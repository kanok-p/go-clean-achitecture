package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	notfound            = "Notfound"
	internalServerError = "InternalServerError"
	badRequest          = "BadRequest"
	validate            = "Validate"
)

type APIError struct {
	Type string `json:"type"`
	Err  error  `json:"error"`
}

func (e APIError) Error() string {
	return e.Err.Error()
}

type ErrorResponse struct {
	Error string      `json:"error"`
	Msg   interface{} `json:"msg"`
}

type Success struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func Notfound(err error) *APIError {
	return &APIError{
		Type: notfound,
		Err:  err,
	}
}

func InternalServerError(err error) *APIError {
	return &APIError{
		Type: internalServerError,
		Err:  err,
	}
}

func BadRequest(err error) *APIError {
	return &APIError{
		Type: badRequest,
		Err:  err,
	}
}

func Validate(err error) *APIError {
	return &APIError{
		Type: validate,
		Err:  err,
	}
}

func OK(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusOK, &Success{
		Code: strconv.Itoa(http.StatusOK),
		Data: resp,
	})
}

func Created(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusCreated, &Success{
		Code: strconv.Itoa(http.StatusCreated),
		Data: resp,
	})
}

func Error(ctx *gin.Context, err error) {
	switch err.(type) {
	case *APIError:
		customError(ctx, err)
	default:
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Error: strconv.Itoa(http.StatusInternalServerError),
			Msg:   err.Error(),
		})
		break
	}

}

func customError(ctx *gin.Context, err error) {
	apiError := err.(*APIError)

	switch apiError.Type {
	case notfound:
		ctx.JSON(http.StatusNotFound, &ErrorResponse{
			Error: strconv.Itoa(http.StatusNotFound),
			Msg:   apiError.Err.Error(),
		})
	case badRequest:
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Error: strconv.Itoa(http.StatusBadRequest),
			Msg:   apiError.Err.Error(),
		})
	case validate:
		ctx.JSON(http.StatusUnprocessableEntity, &ErrorResponse{
			Error: strconv.Itoa(http.StatusUnprocessableEntity),
			Msg:   apiError.Err.Error(),
		})
	case internalServerError:
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Error: strconv.Itoa(http.StatusInternalServerError),
			Msg:   apiError.Err.Error(),
		})
	}
}

type ListResp struct {
	Pagination
	List interface{}
}

type Pagination struct {
	Total  int64
	Page   int64
	Limit  int64
	Search string
	Sort   string
}
