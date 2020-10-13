package inout

type User struct {
	CitizenID    string `json:"citizen_id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobile_number"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	BirthDate    string `json:"birth_date"`
	Gender       string `json:"gender"`
}

