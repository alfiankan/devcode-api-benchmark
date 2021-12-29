package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func NewMysqlConnection() *sql.DB {
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DBNAME := os.Getenv("MYSQL_DBNAME")
	//log.Println(MYSQL_HOST, MYSQL_DBNAME, MYSQL_USER, MYSQL_PASSWORD)

	//sentry.CaptureMessage(MYSQL_HOST + MYSQL_DBNAME + MYSQL_USER + MYSQL_PASSWORD)
	db, err := sql.Open("mysql", MYSQL_USER+":"+MYSQL_PASSWORD+"@tcp("+MYSQL_HOST+")/"+MYSQL_DBNAME+"?charset=utf8mb4&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:473550@tcp(127.0.0.1:3306)/devcode?charset=utf8mb4&parseTime=True&loc=Local")

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Minute * 10)
	//sentry.CaptureException(db.Ping())

	if err != nil {
		log.Println(err)
	}
	return db
}
