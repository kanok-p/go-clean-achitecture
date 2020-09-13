package validate

import (
	"context"
	"errors"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"

	"github.com/kanok-p/go-clean-architecture/config"
	domainusers "github.com/kanok-p/go-clean-architecture/domain/users"
	USRMock "github.com/kanok-p/go-clean-architecture/repository/users/mocks"
)

var formats = map[string][]string{
	"password": {`[0-9]`, `[a-z]`, `[A-Z]`, `^[0-9a-zA-Z]+$`},
}

type TestValidate struct {
	CitizenID    string `validate:"required,unique"`
	Email        string `validate:"required,email,unique"`
	Password     string `validate:"required,password-format,min=8,max=256"`
	MobileNumber string `validate:"required,unique"`
	BirthDate    string `validate:"omitempty,date-format=2006-01-02"`
}

var testValidateSuccess = TestValidate{
	CitizenID:    "1234567890123",
	Email:        "test123@gmail.com",
	Password:     "Te5tSuccess",
	MobileNumber: "0888008800",
	BirthDate:    "1970-01-01",
}

var testValidateError = TestValidate{
	CitizenID:    "1234567890123",
	Email:        "test123@gmail",
	Password:     "testfail",
	MobileNumber: "0888008800",
	BirthDate:    "1970-01-01111",
}

type VLDSuite struct {
	suite.Suite
	usrRepo   *USRMock.Repository
	config    *config.Config
	formats   map[string][]string
	validator Validator
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(VLDSuite))
}

func (suite *VLDSuite) SetupTest() {
	_ = godotenv.Load("/Users/pop/github/go-clean-architecture/.env")
	conf, err := config.Get()
	suite.NoError(err)
	suite.usrRepo = &USRMock.Repository{}
	suite.validator = New(suite.usrRepo, conf, formats)
}

func (suite *VLDSuite) TestNew() {
	_ = godotenv.Load("/Users/pop/github/go-clean-architecture/.env")
	conf, err := config.Get()
	suite.NoError(err)

	suite.validator = New(suite.usrRepo, conf, formats)
	suite.NotEqual(nil, suite.validator)
}

func (suite *VLDSuite) TestValidateSuccess() {
	suite.usrRepo.On("Get", context.Background(), "CitizenID", testValidateSuccess.CitizenID).Once().Return(&domainusers.Users{}, errors.New("no documents in mongo"))
	suite.usrRepo.On("Get", context.Background(), "Email", testValidateSuccess.Email).Once().Return(&domainusers.Users{}, errors.New("no documents in mongo"))
	suite.usrRepo.On("Get", context.Background(), "MobileNumber", testValidateSuccess.MobileNumber).Once().Return(&domainusers.Users{}, errors.New("no documents in mongo"))

	err := suite.validator.Validate(context.Background(), testValidateSuccess)
	suite.NoError(err)
}

func (suite *VLDSuite) TestValidateError() {
	suite.usrRepo.On("Get", context.Background(), "CitizenID", testValidateSuccess.CitizenID).Once().Return(&domainusers.Users{}, nil)
	suite.usrRepo.On("Get", context.Background(), "Email", testValidateSuccess.Email).Once().Return(&domainusers.Users{}, nil)
	suite.usrRepo.On("Get", context.Background(), "MobileNumber", testValidateSuccess.MobileNumber).Once().Return(&domainusers.Users{}, nil)
	err := suite.validator.Validate(context.Background(), testValidateError)
	suite.Error(err)
}
