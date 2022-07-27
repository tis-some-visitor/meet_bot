package dbfellow

import (
	"database/sql"
	"fmt"
	"log"

	"meet_bot/envar"
	"meet_bot/evil"
	"os"

	_ "github.com/lib/pq"
)

type DBConnCreds struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

var Db *sql.DB

//writes DBConnCreds into a string formatted for sql.Open
func createDBConnCredentials() string {
	var host = os.Getenv("host")
	var portDB = os.Getenv("portDB")
	var user = os.Getenv("user")
	var password = os.Getenv("password")
	var dbname = os.Getenv("dbname")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portDB, user, password, dbname)
}

//opens a database and verifies the connection in process
//pass "local" for access to local DB, "heroku" heroku.com
func InitDB(creds string) {

	var err error
	envar.LoadEnvironmentals()

	if creds == "local" {

		cr := createDBConnCredentials()
		Db, err = sql.Open("postgres", cr)
		if err != nil {
			log.Fatal(err)
		}
	} else if creds == "db" {

		var URI = os.Getenv("dbURI")
		Db, err = sql.Open("postgres", URI)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = Db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = evil.RetryPostgres(createTables)
	if err != nil {
		log.Fatal(err)
	}
}
