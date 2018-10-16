package main

import (
	"log"

	"./lib/database/dblayer"
	"./orgpa"
)

func main() {
	databaseHandler, err := dblayer.NewDBLayer(dblayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = orgpa.Run(databaseHandler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
