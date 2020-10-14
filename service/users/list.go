package users

import (
	"context"
	"reflect"

	"github.com/iancoleman/strcase"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kanok-p/go-clean-architecture/domain/request"
	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (u USRService) List(ctx context.Context, input *request.PageOption) (int64, []*domain.Users, error) {
	filter := makeFilter(domain.Users{}, input.Search)
	total, list, err := u.usrRepo.List(ctx, makeOffset(input.Page, input.PerPage), input.PerPage, filter)
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
		keys = append(keys, strcase.ToLowerCamel(userStruct.Type().Field(i).Name))
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

func makeOffset(page, perPage int64) (skip int64) {
	return (page - 1) * perPage
}
