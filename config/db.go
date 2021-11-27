package config

import (
	"fmt"
	"log"
	"os"

	"api-golearn/v1/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASS")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, Password, Host, Port, DbName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection failed:", err)

	}
	DB.AutoMigrate(entity.User{})

	fmt.Println("DB Connected")
	return DB
}
