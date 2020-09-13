package app

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/kanok-p/go-clean-architecture/app/inout"
	USRMock "github.com/kanok-p/go-clean-architecture/service/users/mocks"
)

const (
	CitizenID    = "1234567890123"
	Email        = "test@mail.com"
	Password     = "Pa55w0rd"
	MobileNumber = "0888808800"
	FirstName    = "FirstName"
	LastName     = "LastName"
	BirthDate    = "1970-01-01"
	Gender       = "male"
)

type AppTestSuite struct {
	suite.Suite
	usersService *USRMock.Service

	app    *App
	router *gin.Engine
}

func (s *AppTestSuite) SetupTest() {
	s.usersService = &USRMock.Service{}

	app := New(s.usersService)

	gin.SetMode("release")
	router := gin.New()

	app.Register(router)

	s.app = app
	s.router = router
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}

var input = inout.User{
	CitizenID:    CitizenID,
	Email:        Email,
	Password:     Password,
	MobileNumber: MobileNumber,
	FirstName:    FirstName,
	LastName:     LastName,
	BirthDate:    BirthDate,
	Gender:       Gender,
}
