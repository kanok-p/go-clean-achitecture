package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	domainUsr "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (s Store) Get(ctx context.Context, field string, value interface{}) (users *domainUsr.Users, err error) {
	err = s.collection().FindOne(ctx, bson.M{field: value}).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
