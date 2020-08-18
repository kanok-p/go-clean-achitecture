package users

import (
	"context"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (u USRService) List(ctx context.Context, offset, limit int64) (int64, []*domainUsr.Users, error) {
	total, list, err := u.usrRepo.List(ctx, offset, limit)
	if err != nil {
		return total, nil, err
	}

	return total, list, nil
}
