package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Store) Delete(ctx context.Context, input *primitive.ObjectID) error {
	_, err := s.collection().DeleteOne(ctx, bson.M{_ID: input})
	if err != nil {
		return err
	}

	return nil
}
