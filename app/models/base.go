package models

import(
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Ticket struct {
	ID			uint 	`json:"id" gorm:"primary_key"`
	Title 		string 	`json:"title"`
	Description	string 	`json:"description"`
}

var DB *gorm.DB

func ConnectDB() {
	var err error

	db, err := gorm.Open("sqlite3", "focus_dev")
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Ticket{})

	DB = db
}