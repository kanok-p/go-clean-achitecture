package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (s *Store) List(ctx context.Context, offset, limit int64) (int64, []*domainUsr.Users, error) {
	total, err := s.collection().CountDocuments(ctx, bson.M{})
	if err != nil {
		return total, nil, err
	}

	cursor, err := s.collection().Find(ctx, bson.M{}, options.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		return total, nil, err
	}
	defer cursor.Close(ctx)

	list := make([]*domainUsr.Users, 0)
	for cursor.Next(ctx) {
		user := &domainUsr.Users{}
		if err := cursor.Decode(user); err != nil {
			return total, nil, err
		}

		list = append(list, user)
	}

	return total, list, nil
}
