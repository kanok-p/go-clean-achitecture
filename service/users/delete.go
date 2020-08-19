package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-achitecture/domain/response"
	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (u USRService) Delete(ctx context.Context, input string) (users *domainUsr.Users, err error) {
	id, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return nil, response.BadRequest(err)
	}

	users, err = u.usrRepo.Get(ctx, &id)
	if err != nil {
		return nil, response.Notfound(err)
	}

	err = u.usrRepo.Delete(ctx, &id)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	return users, nil
}
