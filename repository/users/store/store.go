package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/kanok-p/go-clean-architecture/config"
)

const (
	_ID = "_id"
)

type Store struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

func New(config *config.Config) (*Store, error) {
	clientOptions := options.Client().ApplyURI(config.MongoDBEndpoint)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &Store{
		dbName:         config.MongoDBName,
		collectionName: config.MongoDBCollUser,
		db:             db,
	}, nil
}

func (s *Store) collection() *mongo.Collection {
	return s.db.Database(s.dbName).Collection(s.collectionName)
}
