package users

import (
	"context"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domainUsr "github.com/kanok-p/go-clean-architecture/domain/users"
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

func (u USRService) Update(ctx context.Context, input *UpdateUsers) (users *domainUsr.Users, err error) {
	data, err := u.usrRepo.Get(ctx, _ID, input.ID)
	if err != nil {
		return nil, response.Notfound(err)
	}

	users = domainUsr.Update()
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
