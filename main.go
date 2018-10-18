package main

import (
	"fmt"
	"log"

	"github.com/frouioui/orgpa-database-api/configuration"
	"github.com/frouioui/orgpa-database-api/database/dblayer"
	"github.com/frouioui/orgpa-database-api/orgpa"
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
