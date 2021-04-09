package data

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id      int    `gorm:"primaryKey"`
	Name    string `gorm:"name"`
	Surname string `gorm:"surname"`
}

type UserData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{db: db}
}
func (u UserData) ReadAll() ([]User, error) {
	var users []User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	return users, nil
}

func (u UserData) Add(user User) (int, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create user to database, error: %w", result.Error)
	}
	return user.Id, nil
}
