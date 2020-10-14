package inout

import "go.mongodb.org/mongo-driver/bson/primitive"

type Create struct {
	CitizenID    string `json:"citizen_id" validate:"required,unique"`
	Email        string `json:"email" validate:"required,email,unique"`
	Password     string `json:"password" validate:"required,password-format,min=8,max=256"`
	MobileNumber string `json:"mobile_number" validate:"required,unique"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date" validate:"omitempty,date-format=2006-01-02"`
	Gender       string `json:"gender"`
	CreatedAt    *int64
	UpdatedAt    *int64
}

type Update struct {
	ID           *primitive.ObjectID
	CitizenID    string `json:"citizen_id" validate:"required,unique-update"`
	Email        string `json:"email" validate:"required,email,unique-update"`
	MobileNumber string `json:"mobile_number" validate:"required,unique-update"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date" validate:"omitempty,date-format=2006-01-02"`
	Gender       string `json:"gender"`
}
