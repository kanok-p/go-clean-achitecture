package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kanok-p/go-clean-architecture/domain/users"
)

func (s Store) Save(ctx context.Context, input *users.Users) error {
	filter := bson.M{_ID: input.ID}
	update := bson.M{
		"$set": input,
	}
	_, err := s.collection().UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	return err
}
