package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Email       string             `json:"email" bson:"email"`
	Phone       string             `json:"phone" bson:"phone"`
	DateOfBirth string             `json:"dateOfBirth" bson:"dateofbirth"`
}
