package test

import (
	"devcode/internal/database"
	"testing"
)

func TestNewMysqlConnection(t *testing.T) {
	database.NewMysqlConnection()
	//db.Exec("CREATE TABLE IF NOT EXISTS `activity` (`id` int(11) NOT NULL AUTO_INCREMENT,`email` varchar(255) NOT NULL,`title` varchar(255) NOT NULL,`password` varchar(255) NOT NULL,`created_at` datetime NOT NULL,`updated_at` datetime NOT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;")
}

func TestMigrateDB(t *testing.T) {
	// db := database.NewMysqlConnection()
	// database.Migrate(db, "activity")
}
