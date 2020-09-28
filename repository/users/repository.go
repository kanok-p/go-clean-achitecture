package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "github.com/kanok-p/go-clean-architecture/domain/users"
)

//go:generate mockery --name=Repository
type Repository interface {
	List(ctx context.Context, offset, limit int64, filter bson.M) (int64, []*domain.Users, error)
	Get(ctx context.Context, input map[string]interface{}) (users *domain.Users, err error)
	Save(ctx context.Context, input *domain.Users) error
	Delete(ctx context.Context, input *primitive.ObjectID) error
}
