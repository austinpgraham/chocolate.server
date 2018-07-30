
package user

import (
	   "github.com/jinzhu/gorm"
	db "github.com/austinpgraham/chocolate.server/pkg/database"
)

const SESSION_TABLE = "sessions"

type Session struct {
	AssocUser User `gorm:"foreignkey:UserID;association_foreignkey:UserID"`
	UserID uint `gorm:"unique"`
	Cookie string
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
		db.Where("user_id = ?", uid).Delete(Session{})
		err = db.Create(&session).Error
		return err
	}
	return nil
}