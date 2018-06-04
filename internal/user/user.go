package user

import (
	"fmt"

	db "github.com/austinpgraham/chocolate.server/internal/database"
)

const USERS_TABLE = "users"

type User struct {
	UserID uint `json:"id" gorm:"AUTO_INCREMENT;unique_index"`
	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"password,omitempty"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	FByF string `json:"fbyf" gorm:"unique"`
}

func checkTable() {
	db, _ := db.GetConnection()
	defer db.Close()
	if !db.HasTable(USERS_TABLE) {
		db.CreateTable(&User{})
	}
}

func GetUser(att string, val string) *User {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	var user User
	if db.First(&user, fmt.Sprintf("%v = ?", att), val).Error != nil {
		return nil
	}
	return &user
}

func CreateUser(user *User) error {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	err := db.Create(user).Error
	return err
}