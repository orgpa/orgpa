package main

import (
	"fmt"
	"log"
	"orgpa/orgpa-database-api/configuration"
	"orgpa/orgpa-database-api/database/dblayer"
	"orgpa/orgpa-database-api/orgpa"
)

func main() {
	config, err := configuration.ExtractConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(config)
	databaseHandler, err := dblayer.NewDBLayer(config.DBType, config.DBConnection)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = orgpa.Run(databaseHandler, config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
