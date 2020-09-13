package app

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"github.com/kanok-p/go-clean-architecture/app/inout"
	serviceUsr "github.com/kanok-p/go-clean-architecture/service/users"
)

func (s *AppTestSuite) TestAppointment() {

	s.usersService.On("Create", mock.Anything, &serviceUsr.CreateUsers{
		CitizenID:    CitizenID,
		Email:        Email,
		Password:     Password,
		MobileNumber: MobileNumber,
		FirstName:    FirstName,
		LastName:     LastName,
		BirthDate:    BirthDate,
		Gender:       Gender,
	}).Return(func(context.Context, *serviceUsr.CreateUsers) error { return nil })

	req, resp := buildRequestCreateUsers(input)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusCreated, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestCreateUsers(input inout.User) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest("POST", "/users", bytes.NewBuffer(inputBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
