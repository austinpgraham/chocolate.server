package neighborhood

import (
	"fmt"
	"math/rand"

	"github.com/jinzhu/gorm"

	db "github.com/austinpgraham/chocolate.server/pkg/database"
	"github.com/austinpgraham/chocolate.server/pkg/user"
)

const NAME = "name"
const TOKEN_LENGTH = 16
const NEIGHBORHOODS_TABLE = "neighborhoods"

var LETTER_RUNES = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Neighborhood struct {
	gorm.Model
	Name      string    `json:"name" gorm:"unique_index"`
	Admin     user.User `json:"admin" gorm:"foreignkey:AdminID"`
	AdminID   uint      `json:"admin_id"`
	Password  string    `json:"password,omitempty"`
	JoinToken string    `json:"join_token"`
}

func checkTable() {
	db, _ := db.GetConnection()
	defer db.Close()
	if !db.HasTable(NEIGHBORHOODS_TABLE) {
		db.CreateTable(&Neighborhood{})
	}
}

func generateToken(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = LETTER_RUNES[rand.Intn(len(LETTER_RUNES))]
	}
	return string(b)
}

func CreateNeighborhood(neigh *Neighborhood) error {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	token := generateToken(TOKEN_LENGTH)
	neigh.JoinToken = token
	err := db.Create(neigh).Error
	return err
}

func GetNeighborhood(att string, val string) *Neighborhood {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	var neighborhood Neighborhood
	if db.First(&neighborhood, fmt.Sprintf("%v = ?", att), val).Error != nil {
		return nil
	}
	return &neighborhood
}

func GetOwnedNeighborhoods(owner *user.User) []Neighborhood {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	var owned []Neighborhood
	owner.Password = ""
	if db.Where("admin_id = ?", owner.Model.ID).Find(&owned).Error != nil {
		return nil
	}
	modified := owned[:0]
	for _, element := range owned {
		element.Admin = *owner
		element.Password = ""
		modified = append(modified, element)
	}
	return modified
}

func GetAllNeighborhoods() []Neighborhood {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	var all []Neighborhood
	if db.Find(&all).Error != nil {
		return nil
	}
	modified := all[:0]
	for _, element := range all {
		adminUser := user.GetUser("ID", fmt.Sprintf("%d", element.AdminID))
		adminUser.Password = ""
		element.Admin = *adminUser
		element.Password = ""
		modified = append(modified, element)
	}
	return modified
}
