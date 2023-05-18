package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ToDo: Read configuration file, write logs
// https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
/* MongoCN is the DB object connection */
var DatabaseCN = dbConnect("mysql", "root", "741", "default")

/* DBConnect make a remote DB connection, returns a valid connection object */
func dbConnect(dbDriver string, dbUser string, dbPassword string, dbName string) (client *sql.DB) {
	client, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp(localhost:6336)/"+dbName+"?parseTime=true")
	
	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

/* DBConnectionAlive ping the DB, if is reachable return true. */
func DBConnectionAlive() bool {
	err := DatabaseCN.Ping ()

	if err != nil {
		log.Println("Error en ping: " + err.Error())
		log.Println("Check your access password and if your internet connection is still valid.")
		return false
	}

	return true
}
