package users

import (
	"context"

	"github.com/jinzhu/copier"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

type CreateUsers struct {
	CitizenID    string
	Email        string
	MobileNumber string
	FirstName    string
	LastName     string
	BirthDate    string
	Gender       string
}

func (u USRService) Create(ctx context.Context, input *CreateUsers) (err error) {
	users := domainUsr.Create()
	if err = copier.Copy(users, &input); err != nil {
		return err
	}

	return u.usrRepo.Save(ctx, users)

}
