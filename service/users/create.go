package users

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/kanok-p/go-clean-achitecture/domain/response"
	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
	"github.com/kanok-p/go-clean-achitecture/util/password"
)

type CreateUsers struct {
	CitizenID    string `validate:"required,unique"`
	Email        string `validate:"required,email,unique"`
	Password     string `json:"-" validate:"required,password-format,min=8,max=256"`
	MobileNumber string `validate:"required,unique"`
	FirstName    string
	LastName     string
	BirthDate    string `validate:"omitempty,date-format=2006-01-02"`
	Gender       string
	CreatedAt    *int64
	UpdatedAt    *int64
}

func (u USRService) Create(ctx context.Context, input *CreateUsers) (err error) {
	users := domainUsr.Create()
	if err = copier.Copy(users, &input); err != nil {
		return response.InternalServerError(err)
	}
	_ = copier.Copy(input, users)
	if users.Password, err = password.Encrypt(input.Password); err != nil {
		return response.Validate(err)
	}

	if err := u.validator.Validate(ctx, input); err != nil {
		return response.Validate(err)
	}

	return u.usrRepo.Save(ctx, users)

}
