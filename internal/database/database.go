package database

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

const DB_CFG_LOC = "config/db.cfg"

type Config struct {
	Flavor string `json:"flavor"`
	Location string `json:"location"`
}

func readConfig() *Config {
	if _, err := os.Stat(DB_CFG_LOC); os.IsNotExist(err) {
		panic("Database config file has not been defined.")
	}
	raw, _ := ioutil.ReadFile(DB_CFG_LOC)
	cfg := Config{}
	json.Unmarshal(raw, &cfg)
	return &cfg 
}

func GetConnection() (*gorm.DB, error) {
	cfg := readConfig()
	db, err := gorm.Open(cfg.Flavor, cfg.Location)
	if err != nil {
		return nil, err
	}
	return db, nil
}