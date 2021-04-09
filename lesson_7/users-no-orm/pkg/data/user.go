package data

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id int
	Name string
	Surname string
}

type UserData struct {
	db *sql.DB
}

func NewUserData(db *sql.DB) *UserData{
	return &UserData{db: db}
}
func (u UserData) ReadAll() ([]User, error) {
	var users []User
	rows, err := u.db.Query("SELECT * from users")
	if err != nil {
		return nil, fmt.Errorf("can't get users from database, error:%w", err)
	}
	for rows.Next() {
		var tempUser User
		err = rows.Scan(&tempUser.Id, &tempUser.Name, &tempUser.Surname)
		if err != nil {
			return nil, fmt.Errorf("can't scan users from database, error:%w", err)
		}
		users = append(users, tempUser)
	}
	return users, nil
}

func (u UserData) Add(user User) error {
	_, err := u.db.Exec("INSERT INTO users (name, surname) VALUES ($1, $2)", user.Name, user.Surname)
	if err != nil {
		return fmt.Errorf("can't inser user to database, error: %w", err)
	}
	return nil
}
