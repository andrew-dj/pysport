package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sportorg/online/api"
	"github.com/sportorg/online/database"
	"os"
)

var (
	config = flag.String("config", "./data/config.json", "Path to config file")
)

type Configuration struct {
	Addr            string `json:"addr"`
	PathToPWA       string `json:"path_to_pwa"`
	MySQLSourceName string `json:"mysql_source_name"`
}

func main() {
	flag.Parse()
	file, _ := os.Open(*config)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	db := database.Connect(configuration.MySQLSourceName)
	defer db.Close()
	err = db.Ping()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	database.RunMigrations(db)
	fmt.Println("Running Sportorg Online")
	api.Run(configuration.Addr, configuration.PathToPWA)
}
