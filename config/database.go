package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

var DB *sql.DB

func Database() {
	database := "mysql"
	username := "root"
	password := ""
	dbname := "goweb"
	host := "localhost"
	port := 3307

	db, err := sql.Open(database, username+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("database connected")
	DB = db
}
