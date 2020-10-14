package users

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
	"github.com/kanok-p/go-clean-architecture/util/password"
)

func (u USRService) Create(ctx context.Context, input *inout.Create) (err error) {
	users := domain.Create()
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
