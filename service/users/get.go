package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

type test struct {
	A string
	B string
}

func (u USRService) Get(ctx context.Context, input string) (users *domain.Users, err error) {
	id, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return nil, response.BadRequest(err)
	}
	filter := map[string]interface{}{
		_ID: &id,
	}

	users, err = u.usrRepo.Get(ctx, filter)
	if err.Error() == "mongo: no documents in result" {
		return nil, response.Notfound(err)
	}

	return

}

//func MakeFilters(search map[string]interface{}, option string) (filter bson.M) {
//	slFilter := bson.A{}
//	for key, value := range search {
//		slFilter = append(slFilter, bson.M{
//			key: bson.M{
//				"$regex":   value,
//				"$options": "i",
//			},
//		})
//	}
//	operation := fmt.Sprintf("$%s", option)
//	filter = bson.M{
//		operation: slFilter,
//	}
//
//	return filter
//}
