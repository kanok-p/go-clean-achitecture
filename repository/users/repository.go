package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	domainUsr "github.com/kanok-p/go-clean-achitecture/domain/users"
)

type Repository interface {
	List(ctx context.Context, offset, limit int64) (int64, []*domainUsr.Users, error)
	Get(ctx context.Context, input *primitive.ObjectID) (*domainUsr.Users, error)
	Save(ctx context.Context, input *domainUsr.Users) error
	Delete(ctx context.Context, input *primitive.ObjectID) error
}
