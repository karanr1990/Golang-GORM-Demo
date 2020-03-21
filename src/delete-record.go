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
	Id int `gorm:"primary_key";"AUTO_INCREMENT"`
	Name string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)”`
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
	//db.Debug().DropTableIfExists(&UserModel{})

	//Auto create table based on Model
	db.Debug().AutoMigrate(&UserModel{})

	// select records to delete it
	db.Table("user_models").Where("address = ?", "Houston").Delete(&UserModel{})
	
	//Find the record and delete it
        //db.Where("address=?", "Los Angeles").Delete(&UserModel{})

        // Select all records from a model and delete all
        //db.Model(&UserModel{}).Delete(&UserModel{})

	log.Println("delete record")
	
	

}
