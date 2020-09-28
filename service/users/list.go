package users

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/kanok-p/go-clean-architecture/domain/request"
	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (u USRService) List(ctx context.Context, input *request.GetListInput) (int64, []*domain.Users, error) {
	filter := makeFilter(domain.Users{}, input.Search)
	total, list, err := u.usrRepo.List(ctx, input.Offset, input.Limit, filter)
	if err != nil {
		return total, nil, response.InternalServerError(err)
	}

	return total, list, nil
}

func makeFilter(structure interface{}, search string) (filter bson.M) {
	if search == "" {
		return bson.M{}
	}

	var keys []string
	userStruct := reflect.Indirect(reflect.ValueOf(structure))
	for i := 0; i < (userStruct.NumField() - 1); i++ {
		keys = append(keys, userStruct.Type().Field(i).Name)
	}

	slFilter := bson.A{}
	for _, key := range keys {
		slFilter = append(slFilter, bson.M{
			key: bson.M{
				"$regex":   search,
				"$options": "i",
			},
		})
	}

	filter = bson.M{
		"$or": slFilter,
	}

	return filter
}
