package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (s Store) Save(ctx context.Context, input *users.Users) error {
	filter := bson.M{"_id": input.ID}
	update := bson.M{
		"$set": bson.M{
			"CitizenID":    input.CitizenID,
			"Email":        input.Email,
			"MobileNumber": input.MobileNumber,
			"FirstName":    input.FirstName,
			"LastName":     input.LastName,
			"BirthDate":    input.BirthDate,
			"Gender":       input.Gender,
			"CreatedAt":    input.CreatedAt,
			"UpdatedAt":    input.UpdatedAt,
		},
	}

	_, err := s.collection().UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	return err
}
