package neighborhood

import (
	"fmt"

	"github.com/jinzhu/gorm"

	db "github.com/austinpgraham/chocolate.server/pkg/database"
	"github.com/austinpgraham/chocolate.server/pkg/user"
)

const NEIGHBORHOODS_TABLE = "neighborhoods"

type Neighborhood struct {
	gorm.Model
	Name    string    `json:"name" gorm:"unique_index"`
	Admin   user.User `json:"admin" gorm:"foreignkey:AdminID"`
	AdminID uint      `json:"admin_id"`
}

func checkTable() {
	db, _ := db.GetConnection()
	defer db.Close()
	if !db.HasTable(NEIGHBORHOODS_TABLE) {
		db.CreateTable(&Neighborhood{})
	}
}

func CreateNeighborhood(neigh *Neighborhood) error {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
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
