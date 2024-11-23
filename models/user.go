package models

import "startup/config"

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func FindAllUsers() ([]User, error) {
	var users []User
	query := "select id, username, email from users"

	err := config.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
