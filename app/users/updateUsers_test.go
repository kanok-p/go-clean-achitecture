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

	"github.com/kanok-p/go-clean-architecture/app"
	"github.com/kanok-p/go-clean-architecture/app/inout"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
	service "github.com/kanok-p/go-clean-architecture/service/users"
)

func (s *app.AppTestSuite) TestUpdateUsers() {
	objID, err := primitive.ObjectIDFromHex(app.ID)
	s.NoError(err)

	s.usersService.On("Update", mock.Anything, &service.UpdateUsers{
		ID:           &objID,
		CitizenID:    app.CitizenID,
		Email:        app.Email,
		MobileNumber: app.MobileNumber,
		FirstName:    app.FirstName,
		LastName:     app.LastName,
		BirthDate:    app.BirthDate,
		Gender:       app.Gender,
	}).Return(&domain.Users{}, nil)

	req, resp := buildRequestUpdateUsers(&app.input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusOK, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *app.AppTestSuite) TestUpdateUsersError() {
	objID, err := primitive.ObjectIDFromHex(app.ID)
	s.NoError(err)

	s.usersService.On("Update", mock.Anything, &service.UpdateUsers{
		ID:           &objID,
		CitizenID:    app.CitizenID,
		Email:        app.Email,
		MobileNumber: app.MobileNumber,
		FirstName:    app.FirstName,
		LastName:     app.LastName,
		BirthDate:    app.BirthDate,
		Gender:       app.Gender,
	}).Return(&domain.Users{}, errors.New("create_users_error"))

	req, resp := buildRequestUpdateUsers(&app.input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestUpdateUsers(input *inout.User) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/users/%s", app.ID), bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}

func (s *app.AppTestSuite) TestUpdateUsersErrorBadRequest() {
	input := "test_bad_request"
	req, resp := buildRequestUpdateUsersError(app.ID, &input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *app.AppTestSuite) TestUpdateUsersErrorBadRequestID() {
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
