package main

import (
	"fmt"
	"log"
	"orgpa-database-api/configuration"
	"orgpa-database-api/database/dblayer"
	"orgpa-database-api/orgpa"
)

func main() {
	// Extraction of the configuration
	config, err := configuration.ExtractConfiguration()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Creation of a new database layer (connection & interaction with the database)
	databaseHandler, err := dblayer.NewDBLayer(config.DBType, config.DBConnection, config.PasswordMySQL, config.DatabaseName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Server's configuration:", config)

	// Run the main server
	err = orgpa.Run(databaseHandler, config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
