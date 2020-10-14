package users

import (
	"context"

	"github.com/stretchr/testify/mock"

	domain "github.com/kanok-p/go-clean-architecture/domain/users"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (suite *USRSuite) TestCreate() {
	input := inout.Create{}

	suite.validator.On("Validate", mock.Anything, &input).Once().Return(nil)
	suite.repoUsers.On("Save", mock.Anything, domain.Create()).Once().Return(nil)
	err := suite.service.Create(context.Background(), &input)

	suite.NoError(err)
}
