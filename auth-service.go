package main

import (
	"fmt"
	"log"
	"time"

	"github.com/agomezguru/cti-b2b/auth-service/db"
	"github.com/agomezguru/cti-b2b/auth-service/handlers"
)

func main()  {
	// ToDo: Remember to read config file and take vars from him

	// Read yaml configuration file
	// Ref.: https://dev.to/mr_destructive/golang-json-yaml-toml-config-file-reading-22cc
	// Ref.: https://zetcode.com/golang/yaml/
	// Ref.: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
/*
	config, err := os.ReadFile("./settings.yaml")

	if err != nil {
			log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(config, &data)

	if err != nil {
			log.Fatal(err)
	}

	log.Println(data)
	
	for key, value := range data {

			log.Println(key, ": ", value)
	} */

	log.Println("Initializing. Try to connet with DB, please wait...")

	dbConnect(1)

	log.Println("Web Server startup on. Started listening connections.")

	handlers.Drivers()


	fmt.Println(time.Now())
	/*_, finded, userId := db.UserExist("alagunas", "agomez@noticiasnet.mx")

	if finded {
		fmt.Println("Usuario ya existe: " + fmt.Sprint(userId))
	} else {
		fmt.Println("Usuario no regsitrado: " + fmt.Sprint(userId))
	} */
	// All opened conections to database be closed before finish ejecution.
	defer db.DatabaseCN.Close()
}

func dbConnect(try int) {
	i := 1
	if db.DBConnectionAlive() == false && try < 100 {
		
		// Incremet waiting time value
		for i <= try {
			time.Sleep(1000 * time.Millisecond) // 1 second 
			i++
		}
		
		fmt.Printf("DB connection failed. Try again after %d second(s).\n\n", try)
		try++
		dbConnect(try)
	}
	return
}