package main

import (
	"IroChest/config"
	"IroChest/db"
	"IroChest/web"

	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.Info("Initialization begin.")
	// Load configuration.
	config.Init()
	// Initialize the db.
	db.Init()
	web.Init()
	logrus.Info("Initialization end.")
}

func main() {
	Init()
	web.Run()
}
