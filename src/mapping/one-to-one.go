package main

//one to on association
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.ibm.com/Quest-CIO/Golang-Sql-Demo/src/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "golangdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	defer db.Close()

	if err != nil{
		panic(err)
	}

	db.DropTableIfExists(&model.Place{},&model.Town{})

	db.AutoMigrate(&model.Place{},&model.Town{})

	//We need to add foreign keys manually.
	db.Model(&model.Place{}).AddForeignKey("town_id", "towns(id)", "CASCADE", "CASCADE")

	t1 := model.Town{
		Name: "Pune",
	}
	t2 := model.Town{
		Name: "Mumbai",
	}
	t3 := model.Town{
		Name: "Hyderabad",
	}

	p1 := model.Place{
		Name: "Katraj",
		Town: t1,
	}
	p2 := model.Place{
		Name: "Thane",
		Town: t2,
	}
	p3 := model.Place{
		Name: "Secundarabad",
		Town: t3,
	}

	db.Save(&p1)
	db.Save(&p2)
	db.Save(&p3)

	fmt.Println("t1==>", t1, "p1==>", p1)
	fmt.Println("t2==>", t2, "p2s==>", p2)
	fmt.Println("t2==>", t3, "p2s==>", p3)

	//Delete
	db.Where("name=?", "Hyderabad").Delete(&model.Town{})

	//Update
	db.Model(&model.Place{}).Where("id=?", 1).Update("name", "Shivaji Nagar")

	//Select
	places := model.Place{}
	towns := model.Town{}
	fmt.Println("Before Association", places)
	db.Where("name=?", "Shivaji Nagar").Find(&places)
	fmt.Println("After Association", places)
	err = db.Model(&places).Association("town").Find(&places.Town).Error
	fmt.Println("After Association", towns, places)
	fmt.Println("After Association", towns, places, err)

	defer db.Close()

}
