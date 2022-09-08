package config

import (
	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
)

var HOST string = "10.16.1.121"
var PORT string = "8060"
var USERNAME string = "root"
var PASSWORD string = "root"
var DATABASE_NAME string = "sharing_go"

func NewDB() {
	// cfg := mysql.Config{
	// 	User:   USERNAME,
	// 	Passwd: PASSWORD,
	// 	Net:    "tcp",
	// 	Addr:   fmt.Sprintf("%s:%s", HOST, PORT),
	// 	DBName: DATABASE_NAME,
	// }

	// // Get a database handle.
	// _, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
