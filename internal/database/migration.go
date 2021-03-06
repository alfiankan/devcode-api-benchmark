package database

import (
	"database/sql"
	"log"
	"sync"
)

func Migrate(db *sql.DB) {
	//create activity table
	wg := sync.WaitGroup{}
	var err error
	go func() {
		wg.Add(1)
		_, err = db.Exec("DROP TABLE IF EXISTS activities")
		_, err = db.Exec(`CREATE TABLE activities (
			id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
			email varchar(255) NOT NULL,
			title varchar(255) NOT NULL,
			created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at datetime DEFAULT NULL
		  ) ENGINE=InnoDB;`)
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		_, err = db.Exec("DROP TABLE IF EXISTS todos")
		_, err = db.Exec(`CREATE TABLE todos (
			id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title varchar(255) NOT NULL,
			activity_group_id int NOT NULL,
			is_active tinyint(1) NOT NULL,
			priority varchar(55) NOT NULL,
			created_at datetime DEFAULT CURRENT_TIMESTAMP,
			updated_at datetime DEFAULT CURRENT_TIMESTAMP,
			deleted_at datetime DEFAULT NULL
		  ) ENGINE=InnoDB;`)
		wg.Done()
	}()

	wg.Wait()

	//sentry.CaptureException(err)
	if err != nil {
		log.Println(err)
		//sentry.CaptureException(err)
		//sentry.CaptureMessage("MIGRATING FAILED")
	}

}
