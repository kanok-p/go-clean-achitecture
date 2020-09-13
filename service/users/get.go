package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domainUsr "github.com/kanok-p/go-clean-architecture/domain/users"
)

func (u USRService) Get(ctx context.Context, input string) (users *domainUsr.Users, err error) {
	id, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return nil, response.BadRequest(err)
	}

	return u.usrRepo.Get(ctx, _ID, &id)

}
