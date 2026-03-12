package main

import (
	"e4cm/internal/service"
	"flag"
)

func main() {

	var csvPath string
	var jsonPath string
	var dbPath string

	flag.StringVar(&csvPath, "csv", "", "csv file path")
	flag.StringVar(&jsonPath, "json", "", "json file path")
	flag.StringVar(&dbPath, "db", "", "db file path")

	if csvPath == "" || jsonPath == "" || dbPath == "" {
		flag.PrintDefaults()
		return
	}

	flag.Parse()
	service.Migrate(csvPath, jsonPath, dbPath)
}
