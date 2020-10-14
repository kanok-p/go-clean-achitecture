package users

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/kanok-p/go-clean-architecture/domain/response"
	domain "github.com/kanok-p/go-clean-architecture/domain/users"
	"github.com/kanok-p/go-clean-architecture/service/users/inout"
)

func (u USRService) Update(ctx context.Context, input *inout.Update) (users *domain.Users, err error) {
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

	if err := u.validator.Validate(ctx, input); err != nil {
		return nil, response.Validate(err)
	}

	err = u.usrRepo.Save(ctx, users)
	if err != nil {
		return nil, response.InternalServerError(err)
	}

	return users, nil
}
