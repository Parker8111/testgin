package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dataUser = "root"
	dataPass = "Verdes1214"
	dataAddr = "127.0.0.1"
	dataPort = "3306"
	dbname   = "test"
)

func Opendata() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dataUser, dataPass, dataAddr, dataPort, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database open error")
	}
	return db
}
