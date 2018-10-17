package main

import (
	"fmt"
	"log"

	"github.com/frouioui/orgpa/lib/configuration"
	"github.com/frouioui/orgpa/lib/database/dblayer"
	"github.com/frouioui/orgpa/orgpa"
)

func main() {
	config, err := configuration.ExtractConfiguration("configuration.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(config)
	databaseHandler, err := dblayer.NewDBLayer(dblayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = orgpa.Run(databaseHandler, config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
