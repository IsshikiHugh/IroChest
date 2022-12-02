package db

import (
	"errors"
)

// Some fixed environment variable.
var DBName = "IroChest.db"                                  // The db name.
var LocalUserAccessKey = "OnlyLocalCanAccessThroughThisKey" // The local user's access key.

// Define the errors.
var ErrInvalidDBName = errors.New("INVALID_DB_NAME")
var ErrConnectDBFail = errors.New("ERR_CONNECT_DB_FAIL")
var ErrMigrateDBFail = errors.New("ERR_MIGRATE_DB_FAIL")
var ErrInsertRecordFail = errors.New("ERR_INSERT_RECORD_FAIL")
var ErrInvalidAccessKey = errors.New("ERR_INVALID_ACCESS_KEY")

// Define the data structure in records table.
type Records struct {
	ID       uint   `gorm:"primary_key"`
	Category uint   `gorm:"not null"`       // The category idx of the records.
	Data     string `gorm:"not null"`       // The record data.
	Created  int64  `gorm:"autoCreateTime"` // The create time of the record.
}

// Define the data structure in admin data table.
// Note that, all the data use the same ID list.
type AdminData struct {
	ID   uint   `gorm:"primary_key"`
	Type int    `gorm:"not null"` // The type of the admin data
	Data string `gorm:"not null"` // The relative data of the type.
	// Type = 0: not use.
	// Type = 1: the authority data.
	// 		Data refers to the access key.
	// Type = 2: the category data.
	// 		Data refers to the name of the category.
}

var TypeAdminAuth = 1
var TypeAdminCategory = 2

// The note struct.
type Note struct {
	Category string `json:"category"`
	Data     string `json:"data"`
}
