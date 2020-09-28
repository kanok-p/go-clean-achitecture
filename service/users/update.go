package users

import (
	"context"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

type UpdateUsers struct {
	ID           *primitive.ObjectID
	CitizenID    string
	Email        string
	MobileNumber string
	FirstName    string
	LastName     string
	BirthDate    string
	Gender       string
}

func (u USRService) Update(ctx context.Context, input *UpdateUsers) (users *domain.Users, err error) {
	filter := map[string]interface{}{
		_ID: input.ID,
	}
	data, err := u.usrRepo.Get(ctx, filter)
	if err != nil {
		return nil, response.Notfound(err)
	}

	users = domain.Update()
	if err = copier.Copy(users, &input); err != nil {
		return nil, response.BadRequest(err)
	}

	users.CreatedAt = data.CreatedAt
	err = u.usrRepo.Save(ctx, users)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	return users, nil
}
