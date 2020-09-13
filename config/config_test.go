package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

const (
	MongodbEndpoint  = "mongodb://touch:secret@localhost:27018/?authSource=admin"
	MongodbName      = "go-clean"
	MongodbCollUsers = "users"
	TimeZone         = "Asia/Bangkok"
)

func TestGet(t *testing.T) {
	_ = godotenv.Load("/Users/pop/github/go-clean-architecture/.env")
	conf, err := Get()
	fmt.Println(conf.MongoDBCollUser)
	require.NoError(t, err)
	require.Equal(t, MongodbEndpoint, conf.MongoDBEndpoint)
	require.Equal(t, MongodbName, conf.MongoDBName)
	require.Equal(t, MongodbCollUsers, conf.MongoDBCollUser)
	require.Equal(t, TimeZone, conf.TimeZone)
}

func TestGetError(t *testing.T) {
	_ = godotenv.Load("/Users/pop/github/go-clean-architecture/.env")
	_ = os.Unsetenv("MONGODB_NAME")
	conf, err := Get()
	require.Error(t, err)
	require.Equal(t, &Config{
		MongoDBEndpoint: "mongodb://touch:secret@localhost:27018/?authSource=admin",
		MongoDBName:     "",
		MongoDBCollUser: "",
		TimeZone:        "",
	}, conf)
}
