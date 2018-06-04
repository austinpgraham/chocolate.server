package user

import (
	"github.com/jinzhu/gorm"

	db "github.com/austinpgraham/chocolate.server/internal/database"
)

const USERS_TABLE = "users"

type User struct {
	gorm.Model
	ID uint `json:"id"` `gorm:"AUTO_INCREMENT;unique_index"`
	Username string `json:"username"` `gorm:"unique_index"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"` `gorm:"unique"`
}

func checkTable() {
	db, _ := db.GetConnection()
	defer db.Close()
	if !db.HasTable(USERS_TABLE) {
		db.CreateTable(&User{})
	}
}

func CreateUser(username string, password string, first string, last string, email string) *User {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	user := User{Username: username, Password: password, FirstName: first, LastName: last, Email: email}
	db.Create(&user)
	return &user
}