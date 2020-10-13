package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "github.com/kanok-p/go-clean-architecture/domain/users"
	service "github.com/kanok-p/go-clean-architecture/service/users"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (s *UsersTestSuite) TestUpdateUsers() {
	objID, err := primitive.ObjectIDFromHex(ID)
	s.NoError(err)

	s.service.On("Update", mock.Anything, &service.UpdateUsers{
		ID:           &objID,
		CitizenID:    CitizenID,
		Email:        Email,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	}).Return(&domain.Users{}, nil)

	req, resp := buildRequestUpdateUsers(&inout.User{
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

	s.Equal(http.StatusOK, resp.Code)
	s.service.AssertExpectations(s.T())
}

func (s *UsersTestSuite) TestUpdateUsersError() {
	objID, err := primitive.ObjectIDFromHex(ID)
	s.NoError(err)

	s.service.On("Update", mock.Anything, &service.UpdateUsers{
		ID:           &objID,
		CitizenID:    CitizenID,
		Email:        Email,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	}).Return(&domain.Users{}, errors.New("create_users_error"))

	req, resp := buildRequestUpdateUsers(&inout.User{
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

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.service.AssertExpectations(s.T())
}

func buildRequestUpdateUsers(input *inout.User) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/users/%s", ID), bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}

func (s *UsersTestSuite) TestUpdateUsersErrorBadRequest() {
	input := "test_bad_request"
	req, resp := buildRequestUpdateUsersError(ID, &input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.service.AssertExpectations(s.T())
}

func (s *UsersTestSuite) TestUpdateUsersErrorBadRequestID() {
	input := "test_bad_request"
	req, resp := buildRequestUpdateUsersError("ID_TEST_BAD_REQUEST", &input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.service.AssertExpectations(s.T())
}

func buildRequestUpdateUsersError(id string, input *string) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/users/%s", id), bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
