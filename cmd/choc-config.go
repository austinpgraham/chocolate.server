package main

import (
	"os"
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"

	db "github.com/austinpgraham/chocolate.server/pkg/database"
)

func main() {
	// Create the config
	cfg := db.Config{}

	// Define SQL flavor flag
	flag.StringVar(&(cfg.Flavor), "flavor", "sqlite", "Flavor of SQL")
	flag.StringVar(&(cfg.Flavor), "f", "sqlite", "Flabor of SQL (shorthand)")

	// Define database location flag
	flag.StringVar(&(cfg.Location), "location", "", "Location of database")
	flag.StringVar(&(cfg.Location), "l", "", "Location of database (shorthand)")

	flag.Parse()

	// Both args are requred
	if cfg.Flavor == "" || cfg.Location == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Write to json file in config subdirectory
	cfgJson, _ := json.Marshal(cfg)
	err := ioutil.WriteFile(db.DB_CFG_LOC, cfgJson, 0644)
	if err != nil {
		fmt.Printf("Error writing config file.")
	} else {
		fmt.Println("Config file written to", db.DB_CFG_LOC, ".")
	}
}