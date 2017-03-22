package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/kkirsche/echo_cronmon/models"

	"github.com/jinzhu/gorm"
	// SQLite database driver is required
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDB(filepath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filepath)

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

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Host{})
	db.AutoMigrate(&models.Task{})
}
