package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (u USRService) Get(ctx context.Context, input string) (users *domainUsr.Users, err error) {
	id, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return nil, err
	}

	return u.usrRepo.Get(ctx, &id)

}
