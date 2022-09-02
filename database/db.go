package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to DB: %s", err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Successfully connected to DB!")
}
