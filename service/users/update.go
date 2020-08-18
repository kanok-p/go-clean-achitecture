package users

import (
	"context"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
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
	data, err := u.usrRepo.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	users = domainUsr.Update()
	if err = copier.Copy(users, &input); err != nil {
		return nil, err
	}

	users.CreatedAt = data.CreatedAt
	err = u.usrRepo.Save(ctx, users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
