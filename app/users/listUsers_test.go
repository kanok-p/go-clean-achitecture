package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"github.com/kanok-p/go-clean-architecture/domain/request"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

var reqGetListInput = &request.GetListInput{
	Limit:  0,
	Page:   1,
	Offset: 0,
	Search: "",
}

func (s *app.AppTestSuite) TestListUsers() {
	s.usersService.On("List", mock.Anything, &request.GetListInput{}).Return(
		int64(0), []*domain.Users{}, (error)(nil),
	)

	req, resp := buildRequestListUsers(reqGetListInput)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusOK, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func (s *app.AppTestSuite) TestListUsersError() {
	s.usersService.On("List", mock.Anything, &request.GetListInput{}).Return(
		int64(0), nil, errors.New("test_error"),
	)

	req, resp := buildRequestListUsers(reqGetListInput)
	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.usersService.AssertExpectations(s.T())
}

func buildRequestListUsers(reqList *request.GetListInput) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	reqListBytes, _ := json.Marshal(reqList)
	req, _ = http.NewRequest("GET", "/users", bytes.NewBuffer(reqListBytes))
	req.Header.Set("Content-Type", "application/json")

	return req, w
}
