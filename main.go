package main

import (
	"fmt"
	"log"
	"orgpa-database-api/configuration"
	"orgpa-database-api/database/dblayer"
	"orgpa-database-api/orgpa"
)

func main() {
	config, err := configuration.ExtractConfiguration()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Server's configuration:", config)

	databaseHandler, err := dblayer.NewDBLayer(config.DBType, config.DBConnection, config.PasswordMySQL, config.DatabaseName)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = orgpa.Run(databaseHandler, config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
