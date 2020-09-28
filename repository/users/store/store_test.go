package store

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kanok-p/go-clean-architecture/config"
)

type storeUSRSuite struct {
	suite.Suite
	db             *mongo.Client
	dbName         string
	collectionName string
	store          *Store
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(storeUSRSuite))
}

func (suite *storeUSRSuite) SetupTest() {
	_ = godotenv.Load("/Users/pop/go-clean-architecture/.env")
	conf, err := config.Get()
	suite.NoError(err)

	suite.store, err = New(conf)
	suite.NoError(err)
}

func (suite *storeUSRSuite) TestNew() {
	_ = godotenv.Load("/Users/pop/go-clean-architecture/.env")
	conf, err := config.Get()
	suite.NoError(err)

	suite.store, err = New(conf)
	suite.NoError(err)
	suite.NotEqual(nil, suite.store)
}

func (suite *storeUSRSuite) TestNewErrorConnect() {
	_ = godotenv.Load("/Users/pop/go-clean-architecture/.env")
	conf, err := config.Get()
	suite.NoError(err)
	conf.MongoDBEndpoint = conf.MongoDBEndpoint + "with error"

	suite.store, err = New(conf)
	suite.Error(err)
}

func (suite *storeUSRSuite) TestNewErrorPingDatabase() {
	_ = godotenv.Load("/Users/pop/go-clean-architecture/.env")
	conf, err := config.Get()
	suite.NoError(err)
	conf.MongoDBEndpoint = "mongodb://touch:secret1@localhost:27017"

	suite.store, err = New(conf)
	suite.Error(err)
}
