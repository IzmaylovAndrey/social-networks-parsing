package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var DBModels = []interface{} {
	&Users{},
	&Accounts{},
}

func OpenConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=go dbname=go sslmode=disable password=qwerty1234")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate(db gorm.DB){
	if err := db.AutoMigrate(DBModels...).Error; err != nil{
		panic(err)
	}
}

func CloseConnection(db gorm.DB){
	db.Close()
}