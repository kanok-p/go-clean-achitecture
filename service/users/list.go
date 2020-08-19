package users

import (
	"context"

	"github.com/kanok-p/go-clean-achitecture/domain/response"
	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

func (u USRService) List(ctx context.Context, offset, limit int64) (int64, []*domainUsr.Users, error) {
	total, list, err := u.usrRepo.List(ctx, offset, limit)
	if err != nil {
		return total, nil, response.InternalServerError(err)
	}

	return total, list, nil
}
