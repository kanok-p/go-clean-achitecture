package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Type string
	Err  error
}

func (e APIError) Error() string {
	return e.Err.Error()
}

type ErrorResponse struct {
	Error string
	Msg   interface{}
}

type Success struct {
	Code string
	Data interface{}
}

func Notfound(err error) *APIError {
	return &APIError{
		Type: "Notfound",
		Err:  err,
	}
}

func InternalServerError(err error) *APIError {
	return &APIError{
		Type: "InternalServerError",
		Err:  err,
	}
}

func BadRequest(err error) *APIError {
	return &APIError{
		Type: "BadRequest",
		Err:  err,
	}
}

func Validate(err error) *APIError {
	return &APIError{
		Type: "Validate",
		Err:  err,
	}
}

func OK(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusOK, &Success{
		Code: strconv.Itoa(http.StatusOK),
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
	case "Notfound":
		ctx.JSON(http.StatusNotFound, &ErrorResponse{
			Error: strconv.Itoa(http.StatusNotFound),
			Msg:   apiError.Err.Error(),
		})
	case "InternalServerError":
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Error: strconv.Itoa(http.StatusInternalServerError),
			Msg:   apiError.Err.Error(),
		})
	case "BadRequest":
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Error: strconv.Itoa(http.StatusBadRequest),
			Msg:   apiError.Err.Error(),
		})
	case "Validate":
		ctx.JSON(http.StatusUnprocessableEntity, &ErrorResponse{
			Error: strconv.Itoa(http.StatusUnprocessableEntity),
			Msg:   apiError.Err.Error(),
		})
	}
}
