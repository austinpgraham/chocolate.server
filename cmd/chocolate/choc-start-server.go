package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "/Users/austingraham/chocolate-data/metastore.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "L1212", Price: 1000})
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
