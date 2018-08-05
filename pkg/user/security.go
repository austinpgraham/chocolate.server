package user

import (
	"errors"

	db "github.com/austinpgraham/chocolate.server/pkg/database"
	"github.com/jinzhu/gorm"
)

const SESSION_TABLE = "sessions"

type Session struct {
	AssocUser User `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
	UserID    uint `gorm:"unique"`
	Cookie    string
}

func checkSessionTable(db *gorm.DB) {
	if !db.HasTable(SESSION_TABLE) {
		db.CreateTable(&Session{})
	}
}

func SaveSession(uid uint, cookie string) error {
	db, _ := db.GetConnection()
	defer db.Close()
	checkSessionTable(db)
	session := Session{UserID: uid, Cookie: cookie}
	err := db.Create(&session).Error
	if err != nil {
		RemoveSession(uid)
		err = db.Create(&session).Error
		return err
	}
	return nil
}

func RemoveSession(uid uint) {
	db, _ := db.GetConnection()
	defer db.Close()
	checkSessionTable(db)
	db.Where("user_id = ?", uid).Delete(Session{})
}

func GetSession(uid uint) *Session {
	db, _ := db.GetConnection()
	defer db.Close()
	checkSessionTable(db)
	session := Session{}
	db.Where("user_id = ?", uid).First(&session)
	return &session
}

func CheckSession(uid uint, cookie string) error {
	session := GetSession(uid)
	if session.Cookie == cookie {
		return nil
	}
	return errors.New("Session invalid.")
}
