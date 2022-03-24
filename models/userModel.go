package models

type User struct {
	Username    string `json:"username,omitempty" bson:"username" validate:"required"`
	Email       string `json:"email,omitempty" bson:"email" validate:"required"`
	Phone       string `json:"phone,omitempty" bson:"phone" validate:"required"`
	DateOfBirth string `json:"dateOfBirth,omitempty" bson:"dateofbirth" validate:"required"`
}
