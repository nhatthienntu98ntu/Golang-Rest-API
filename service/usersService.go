package service

import (
	"database/sql"
	"fmt"

	"github.com/nhatthienntu98ntu/Golang-Rest-API/config"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/model"
)

func GetAll() ([]model.User, error) {
	db := config.GetDB()
	// An albums slice to hold data from returned rows.
	var users []model.User

	_users, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer _users.Close()
	// Loop through _users, using Scan to assign column data to struct fields.
	for _users.Next() {
		var user model.User
		if err := _users.Scan(&user.Id, &user.Username, &user.Phone, &user.DateOfBirth); err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		users = append(users, user)
	}
	if err := _users.Err(); err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return users, nil
}

func GetUserById(id int) (model.User, error) {
	db := config.GetDB()
	var user model.User

	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Username, &user.Phone, &user.DateOfBirth); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("error: %v, no sunch user", err)
		}
		return user, fmt.Errorf("err: %v", err)
	}

	return user, nil
}

func AddUser(user model.UserInput) (model.User, error) {
	db := config.GetDB()
	var userResponse model.User
	result, err := db.Exec("INSERT INTO users (username, phone, dateOfBirth) VALUE(?, ?, ?)", user.Username, user.Phone, user.DateOfBirth)
	if err != nil {
		return userResponse, fmt.Errorf("error: %v", err)
	}
	fmt.Println(result)
	id, err := result.LastInsertId()
	if err != nil {
		return userResponse, fmt.Errorf("error: %v", err)
	}
	userResponse, err = GetUserById(int(id))
	if err != nil {
		return userResponse, fmt.Errorf("error: %v", err)
	}
	return userResponse, nil
}

func UpdateUser(id int, user model.UserInput) (model.User, error) {
	db := config.GetDB()
	var userResponse model.User
	_, err := db.Exec("UPDATE users SET username = ?, phone = ?, dateOfBirth = ? WHERE id = ? ", user.Username, user.Phone, user.DateOfBirth, id)
	if err != nil {
		return userResponse, fmt.Errorf("error: %v", err)
	}
	userResponse, err = GetUserById(id)
	if err != nil {
		return userResponse, fmt.Errorf("error: %v", err)
	}
	return userResponse, nil
}

func DeleteUser(id int) (int, error) {
	db := config.GetDB()
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}
	return id, nil
}
