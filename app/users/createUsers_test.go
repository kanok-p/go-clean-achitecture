package users

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"github.com/kanok-p/go-clean-architecture/app"
	"github.com/kanok-p/go-clean-architecture/app/inout"
	serviceUsr "github.com/kanok-p/go-clean-architecture/service/users"
)

func (s *app.AppTestSuite) TestCreateUsers() {

	s.usersService.On("Create", mock.Anything, &serviceUsr.CreateUsers{
		CitizenID:    app.CitizenID,
		Email:        app.Email,
		Password:     app.Password,
		MobileNumber: app.MobileNumber,
		FirstName:    app.FirstName,
		LastName:     app.LastName,
		BirthDate:    app.BirthDate,
		Gender:       app.Gender,
	}).Return(func(context.Context, *serviceUsr.CreateUsers) error { return nil })

	req, resp := buildRequestCreateUsers(&app.input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusCreated, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *app.AppTestSuite) TestCreateUsersError() {

	s.usersService.On("Create", mock.Anything, &serviceUsr.CreateUsers{}).Return(
		func(context.Context, *serviceUsr.CreateUsers) error { return errors.New("create_users_error") })

	req, resp := buildRequestCreateUsers(&inout.User{})
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestCreateUsers(input *inout.User) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}

func (s *app.AppTestSuite) TestCreateUsersErrorBadRequest() {
	input := "test_bad_request"
	req, resp := buildRequestCreateUsersError(&input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestCreateUsersError(input *string) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
