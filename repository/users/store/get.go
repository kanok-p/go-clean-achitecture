package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (s Store) Get(ctx context.Context, input *primitive.ObjectID) (users *users.Users, err error) {
	err = s.collection().FindOne(ctx, bson.M{"_id": input}).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
