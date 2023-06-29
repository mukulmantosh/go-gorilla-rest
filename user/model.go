package user

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

const DbUrl = "root:test123@tcp(127.0.0.1:3306)/mukuldb?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitMigration() {
	DB, err = gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("Cannot connect to DB")
	}
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Println(err.Error())
		return
	}

}
