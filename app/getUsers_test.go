package app

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (s *AppTestSuite) TestGetUsers() {
	s.usersService.On("Get", mock.Anything, ID).Return(
		&domain.Users{}, (error)(nil),
	)

	req, resp := buildRequestGetUsers(ID)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusOK, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *AppTestSuite) TestGetUsersError() {
	s.usersService.On("Get", mock.Anything, ID).Return(
		nil, errors.New("test_error"),
	)

	req, resp := buildRequestGetUsers(ID)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestGetUsers(id string) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	req, _ = http.NewRequest("GET", fmt.Sprintf("/users/%s", id), bytes.NewBuffer(nil))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
