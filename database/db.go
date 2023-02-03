package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mk1589/Week-Two-GL-CipherSchools/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "manish@1315"

	args := "host=" + host + "port=" + port + " user=" + username + "dbname=" + dbName + " sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}
