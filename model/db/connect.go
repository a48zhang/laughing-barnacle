package db

import (
	"encoding/json"
	"io"
	"log"
	"os"

	gomysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	file, err := os.Open("./conf/db.json")
	var cert gomysql.Config
	if err != nil {
		cert = gomysql.Config{
			User:                 os.Getenv("DB_USER"),
			Passwd:               os.Getenv("DB_PWD"),
			Net:                  "tcp",
			Addr:                 os.Getenv("DB_ADDR"),
			DBName:               os.Getenv("DB_NAME"),
			AllowNativePasswords: true,
		}
	} else {
		tmp, _ := io.ReadAll(file)
		json.Unmarshal(tmp, &cert)
	}
	DB, err = gorm.Open(mysql.Open(cert.FormatDSN()))
	if err != nil {
		log.Fatal(err)
	}

}
