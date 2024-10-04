package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect establishes a connection to the database.
func Connect(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	connectionStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
	return gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
}
