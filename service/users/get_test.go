package users

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (suite *USRSuite) TestGet() {
	id := primitive.NewObjectID()
	filter := map[string]interface{}{
		_ID: id,
	}
	expectedUser := domain.Create()
	suite.repoUsers.On("Get", mock.Anything, filter).Once().Return(expectedUser, nil)
	users, err := suite.service.Get(context.Background(), id.Hex())

	suite.NoError(err)
	suite.Equal(expectedUser, users)
}

func (suite *USRSuite) TestGetIDError() {
	id := "1234"
	users, err := suite.service.Get(context.Background(), id)

	suite.Error(err)
	suite.Nil(users)
}

func (suite *USRSuite) TestGetErrorNoTFound() {
	id := primitive.NewObjectID()
	filter := map[string]interface{}{
		_ID: id,
	}

	suite.repoUsers.On("Get", mock.Anything, filter).Once().Return(nil, errors.New("mongo: no documents in result"))
	users, err := suite.service.Get(context.Background(), id.Hex())

	suite.Error(err)
	suite.Nil(users)
}
