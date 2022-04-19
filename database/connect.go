package database

import (
	_ "database/sql"
	"fmt"
	"github.com/yeongsummer/restful-api-with-go/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MakeDSN(dbUser string, dbPassword string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
}

func ConnetDB(dsn string) {
	var err error

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")

}
