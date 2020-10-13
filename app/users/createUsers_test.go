package users

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	serviceUsr "github.com/kanok-p/go-clean-architecture/service/users"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (s *UsersTestSuite) TestCreateUsers() {
	s.service.On("Create", mock.Anything, &serviceUsr.CreateUsers{
		CitizenID:    CitizenID,
		Email:        Email,
		Password:     Password,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	}).Return(func(context.Context, *serviceUsr.CreateUsers) error { return nil })

	req, resp := buildRequestCreateUsers(&inout.User{
		CitizenID:    CitizenID,
		Email:        Email,
		Password:     Password,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	})
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusCreated, resp.Code)
	s.service.AssertExpectations(s.T())
}

func (s *UsersTestSuite) TestCreateUsersError() {

	s.service.On("Create", mock.Anything, &serviceUsr.CreateUsers{}).Return(
		func(context.Context, *serviceUsr.CreateUsers) error { return errors.New("create_users_error") })

	req, resp := buildRequestCreateUsers(&inout.User{})
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.service.AssertExpectations(s.T())
}

func buildRequestCreateUsers(input *inout.User) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}

func (s *UsersTestSuite) TestCreateUsersErrorBadRequest() {
	input := "test_bad_request"
	req, resp := buildRequestCreateUsersError(&input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.service.AssertExpectations(s.T())
}

func buildRequestCreateUsersError(input *string) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
