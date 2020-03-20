package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "golangdb"
)

type User struct {
	ID int
	Username string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	//log.Println("Connection Established")

		// After db connection is created.
		db.CreateTable(&User{})

		// Also some useful functions
		db.HasTable(&User{}) // =>;; true
		db.DropTableIfExists(&User{})  //Drops the table if already exists
	log.Println("Connection Established")
}
