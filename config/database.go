package config

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var HOST string = "10.16.1.121"
var PORT string = "8060"
var USERNAME string = "root"
var PASSWORD string = "root"
var DATABASE_NAME string = "sharing_go"

func NewDB() *sql.DB {
	// cfg := mysql.Config{
	// 	User:   USERNAME,
	// 	Passwd: PASSWORD,
	// 	Net:    "tcp",
	// 	Addr:   fmt.Sprintf("%s:%s", HOST, PORT),
	// 	DBName: DATABASE_NAME,
	// }

	// DSN for MYSQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USERNAME, PASSWORD, HOST, PORT, DATABASE_NAME) // "username:password@tcp(127.0.0.1:3306)/jazzrecords"

	// Create Connection to Database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB CONNECTED")
	return db
}

func NewGormDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USERNAME, PASSWORD, HOST, PORT, DATABASE_NAME) // "username:password@tcp(127.0.0.1:3306)/jazzrecords"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
