package users

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"github.com/kanok-p/go-clean-architecture/app"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (s *app.AppTestSuite) TestGetUsers() {
	s.usersService.On("Get", mock.Anything, app.ID).Return(
		&domain.Users{}, (error)(nil),
	)

	req, resp := buildRequestGetUsers(app.ID)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusOK, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *app.AppTestSuite) TestGetUsersError() {
	s.usersService.On("Get", mock.Anything, app.ID).Return(
		nil, errors.New("test_error"),
	)

	req, resp := buildRequestGetUsers(app.ID)
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
