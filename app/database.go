package app

import (
	"database/sql"

	"github.com/Sirupsen/logrus"
	// SQLite database driver is required
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		logrus.WithError(err).Errorln("Failed to connect to database")
		return nil
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		logrus.Errorln("No database connection established, exiting")
		return nil
	}

	return db
}

func migrate(db *sql.DB) error {
	migrateHost(db)
	// if err != nil {
	// 	logrus.WithError(err).Errorln("Failed to migrate host object")
	// 	return err
	// }

	migrateTask(db)
	// if err != nil {
	// 	logrus.WithError(err).Errorln("Failed to migrate task object")
	// 	return err
	// }

	return nil
}
