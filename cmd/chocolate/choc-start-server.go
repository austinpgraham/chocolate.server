package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"github.com/gorilla/mux"

	db "github.com/austinpgraham/chocolate.server/internal/database"
)

func checkDatabaseConnection() error {
	_db, err := db.GetConnection()
	defer _db.Close()
	if err == nil {
		return err
	}
	return nil
}

var port string

func main() {
	// Build required command line arguments
	flag.StringVar(&port, "port", "8000", "Port to run server")
	flag.StringVar(&port, "p", "8000", "Port to run server (shorthane)")
	flag.Parse()

	// Check that we can connect to the database
	fmt.Println("Checking database connection...")
	err := checkDatabaseConnection()
	if err != nil {
		panic("Database connection could not be obtained.")
	}

	// Start the server
	fmt.Println("Starting server on port", port, "...")
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
