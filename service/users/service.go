package users

import (
	"context"

	"github.com/kanok-p/go-clean-achitecture/domain/request"
	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
	"github.com/kanok-p/go-clean-achitecture/repository/users"
	"github.com/kanok-p/go-clean-achitecture/util/validate"
)

const (
	_ID = "_id"
)

type Service interface {
	List(ctx context.Context, input *request.GetListInput) (int64, []*domainUsr.Users, error)
	Create(ctx context.Context, input *CreateUsers) error
	Get(ctx context.Context, input string) (*domainUsr.Users, error)
	Update(ctx context.Context, input *UpdateUsers) (*domainUsr.Users, error)
	Delete(ctx context.Context, input string) (*domainUsr.Users, error)
}

type USRService struct {
	usrRepo   users.Repository
	validator validate.Validator
}

func New(usrRepo users.Repository, validator validate.Validator) *USRService {
	return &USRService{
		usrRepo:   usrRepo,
		validator: validator,
	}
}
