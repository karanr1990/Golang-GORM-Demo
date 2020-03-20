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

type UserModel struct{
	Id int `gorm:"primary_key"`
	Name string
	Address string
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

	log.Println("Connection Established")

	//Drops table if already exists
	db.Debug().DropTableIfExists(&UserModel{})

	//Auto create table based on Model
	db.Debug().AutoMigrate(&UserModel{})

}
