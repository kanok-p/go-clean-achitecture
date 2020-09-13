package users

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/kanok-p/go-clean-architecture/domain/request"
	"github.com/kanok-p/go-clean-architecture/domain/response"
	domainUsr "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (u USRService) List(ctx context.Context, input *request.GetListInput) (int64, []*domainUsr.Users, error) {
	var keys []string
	userStruct := reflect.Indirect(reflect.ValueOf(domainUsr.Users{}))
	for i := 0; i < (userStruct.NumField() - 1); i++ {
		keys = append(keys, userStruct.Type().Field(i).Name)
	}

	filter := makeFilter(keys, input.Search)
	total, list, err := u.usrRepo.List(ctx, input.Offset, input.Limit, filter)
	if err != nil {
		return total, nil, response.InternalServerError(err)
	}

	return total, list, nil
}

func makeFilter(keys []string, search string) (filter bson.M) {
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
