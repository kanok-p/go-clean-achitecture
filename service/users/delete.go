package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (u USRService) Delete(ctx context.Context, input string) (users *domain.Users, err error) {
	id, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return nil, response.BadRequest(err)
	}

	filter := map[string]interface{}{
		_ID: &id,
	}
	users, err = u.usrRepo.Get(ctx, filter)
	if err != nil {
		return nil, response.Notfound(err)
	}

	err = u.usrRepo.Delete(ctx, &id)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	return users, nil
}
