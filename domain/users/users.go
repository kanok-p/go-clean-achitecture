package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	CitizenID    string             `bson:"citizenID" json:"citizen_id"`
	Email        string             `bson:"email" json:"email"`
	Password     string             `bson:"password" json:"-"`
	MobileNumber string             `bson:"mobileNumber" json:"mobile_number"`
	FirstName    string             `bson:"firstName" json:"first_name"`
	LastName     string             `bson:"lastName" json:"last_name"`
	BirthDate    string             `bson:"birthDate" json:"birth_date"`
	Gender       string             `bson:"gender" json:"gender"`
	CreatedAt    int64              `bson:"createdAt" json:"created_at"`
	UpdatedAt    int64              `bson:"updatedAt" json:"updated_at"`
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
