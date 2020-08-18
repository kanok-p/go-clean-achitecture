package users

import (
	"context"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
	"github.com/kanok-p/go-clean-achitecture/repository/users"
)

type Service interface {
	List(ctx context.Context, offset, limit int64) (int64, []*domainUsr.Users, error)
	Create(ctx context.Context, input *CreateUsers) error
	Get(ctx context.Context, input string) (*domainUsr.Users, error)
	Update(ctx context.Context, input *UpdateUsers) (*domainUsr.Users, error)
	Delete(ctx context.Context, input string) (*domainUsr.Users, error)
}

type USRService struct {
	usrRepo users.Repository
}

func New(usrRepo users.Repository) *USRService {
	return &USRService{
		usrRepo: usrRepo,
	}
}
