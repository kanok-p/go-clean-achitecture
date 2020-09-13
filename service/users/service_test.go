package users

import (
	"testing"

	"github.com/stretchr/testify/suite"

	USRMock "github.com/kanok-p/go-clean-architecture/repository/users/mocks"
	VLDMock "github.com/kanok-p/go-clean-architecture/util/validate/mocks"
)

type USRSuite struct {
	suite.Suite
	repoUsers *USRMock.Repository
	validator *VLDMock.Validator
	service   *USRService
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(USRSuite))
}

func (suite *USRSuite) SetupTest() {
	suite.repoUsers = &USRMock.Repository{}
	suite.validator = &VLDMock.Validator{}
	suite.service = New(suite.repoUsers, suite.validator)
}

func (suite *USRSuite) TestNew() {
	suite.service = New(suite.repoUsers, suite.validator)
	suite.NotEqual(nil, suite.service.usrRepo)
	suite.NotEqual(nil, suite.service.validator)
}
