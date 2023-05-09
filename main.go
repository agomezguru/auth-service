package main

import (
	"log"
	"os"

	"github.com/agomezguru/cti-b2b/auth-service/db"

	yaml "gopkg.in/yaml.v3"
)

func main()  {
	// ToDo: Remember to read config file and take vars from him

	// Read yaml configuration file
	// Ref.: https://dev.to/mr_destructive/golang-json-yaml-toml-config-file-reading-22cc
	// Ref.: https://zetcode.com/golang/yaml/
	// Ref.: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
	config, err := os.ReadFile("settings.yaml")

	if err != nil {
			log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(config, &data)

	if err != nil {
			log.Fatal(err)
	}
/*
	log.Println(data)
	
	for key, value := range data {

			log.Println(key, ": ", value)
	} */

	log.Println("Initializing. Try to connet with DB, please wait...")

	if db.DBConnectionAlive() == false {
		log.Fatal("DB connection failed")
		return 
	}
	// All opened conections to database be closed before finish ejecution.
	defer db.DatabaseCN.Close()
}