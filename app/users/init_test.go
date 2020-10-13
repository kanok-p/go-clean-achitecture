package users

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/kanok-p/go-clean-architecture/service/users/mocks"
)

const (
	ID           = "5f8591a510ac5b73eef104bf"
	CitizenID    = "CitizenID"
	Email        = "Email"
	Password     = "Password"
	MobileNumber = "MobileNumber"
	FirstName    = "FirstName"
	LastName     = "LastName"
	BirthDate    = "BirthDate"
	Gender       = "Gender"
)

type UsersTestSuite struct {
	suite.Suite
	router     *gin.Engine
	ctx        *gin.Context
	controller *Controller
	service    *mocks.Service
}

func (s *UsersTestSuite) SetupSuite() {
	s.service = &mocks.Service{}
	s.controller = New(s.service)
	s.ctx, s.router = gin.CreateTestContext(httptest.NewRecorder())

	s.router.GET("/users", s.controller.ListUsers)
	s.router.POST("/users", s.controller.CreateUsers)
	s.router.GET("/users/:id", s.controller.GetUsers)
	s.router.PUT("/users/:id", s.controller.UpdateUsers)
	s.router.DELETE("/users/:id", s.controller.DeleteUsers)
}

func TestFormTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
