package model

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"dateOfBirth"`
}

type UserInput struct {
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"dateOfBirth"`
}
