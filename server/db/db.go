package db

import (
	"IroChest/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialization of the db module.
func Init() {
	logrus.Info("Trying to initialize db module...")
	db = Connect(DBName)
	err := db.AutoMigrate(&Records{}, &AdminData{})
	if err != nil {
		logrus.Fatal(ErrMigrateDBFail.Error())
	}

	// Initialize the local user, i.e., the first user.
	if !QueryHasAccessKey(config.Conf.OriginAccessKeys) {
		AddAccessKey(config.Conf.OriginAccessKeys)
	}
	logrus.Info("Initialized!")
}

// Connect to the sqlite db.
func Connect(name string) *gorm.DB {
	logrus.Info(fmt.Sprintf("Trying to connect %s...", name))
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})

	if err != nil {
		logrus.Fatal(ErrConnectDBFail.Error() + " : " + err.Error())
		return nil
	} else {
		logrus.Info("Success!")
		return db
	}
}
