package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID           primitive.ObjectID `bson:"_id"`
	CitizenID    string
	Email        string
	Password     string `json:"-"`
	MobileNumber string
	FirstName    string
	LastName     string
	BirthDate    string
	Gender       string
	CreatedAt    int64
	UpdatedAt    int64
}

func Create() *Users {
	now := time.Now()
	return &Users{
		ID:        primitive.NewObjectID(),
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}
}

func Update() *Users {
	now := time.Now()
	return &Users{
		UpdatedAt: now.Unix(),
	}
}
