package database

import (
	"fmt"
	"goapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"

const DB_PASSWORD = ""
const DB_NAME = "stock"
const DB_HOST = ""
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error = ", err)
		return nil
	}
	db.AutoMigrate(&models.UserInfo{})
	db.AutoMigrate(&models.Portfolio{})

	return db
}
