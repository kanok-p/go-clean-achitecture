package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

const (
	and = "$and"
)

func (s Store) Get(ctx context.Context, input map[string]interface{}) (users *domain.Users, err error) {
	var slFilter []bson.M
	for key, value := range input {
		slFilter = append(slFilter, bson.M{
			key: value,
		})
	}

	err = s.collection().FindOne(ctx, bson.M{and: slFilter}).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
