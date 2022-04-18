package database

import (
	_ "database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MakeDSN(dbUser string, dbPassword string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
}

func ConnetDB(dsn string) {
	var err error

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

}
