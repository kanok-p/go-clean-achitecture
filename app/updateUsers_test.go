package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/app/inout"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
	service "github.com/kanok-p/go-clean-architecture/service/users"
)

func (s *AppTestSuite) TestUpdateUsers() {
	objID, err := primitive.ObjectIDFromHex(ID)
	s.NoError(err)

	s.usersService.On("Update", mock.Anything, &service.UpdateUsers{
		ID:           &objID,
		CitizenID:    CitizenID,
		Email:        Email,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	}).Return(&domain.Users{}, nil)

	req, resp := buildRequestUpdateUsers(&input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusOK, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *AppTestSuite) TestUpdateUsersError() {
	objID, err := primitive.ObjectIDFromHex(ID)
	s.NoError(err)

	s.usersService.On("Update", mock.Anything, &service.UpdateUsers{
		ID:           &objID,
		CitizenID:    CitizenID,
		Email:        Email,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	}).Return(&domain.Users{}, errors.New("create_users_error"))

	req, resp := buildRequestUpdateUsers(&input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestUpdateUsers(input *inout.User) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/users/%s", ID), bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}

func (s *AppTestSuite) TestUpdateUsersErrorBadRequest() {
	input := "test_bad_request"
	req, resp := buildRequestUpdateUsersError(ID, &input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *AppTestSuite) TestUpdateUsersErrorBadRequestID() {
	input := "test_bad_request"
	req, resp := buildRequestUpdateUsersError("ID_TEST_BAD_REQUEST", &input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestUpdateUsersError(id string, input *string) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/users/%s", id), bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
