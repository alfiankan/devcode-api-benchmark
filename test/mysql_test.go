package test

import (
	"devcode/internal/database"
	"log"
	"testing"
)

func TestNewMysqlConnection(t *testing.T) {
	db := database.NewMysqlConnection()
	log.Println(db.Ping())
}

func TestMigrateDB(t *testing.T) {
	db := database.NewMysqlConnection()
	database.Migrate(db)
}
