package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * `db/admin.go`												 *
 * Be used to do the CURD option for IroChest's categories and   *
 * authority.    												 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// Insert an access key.
func AddAccessKey(key string) error {
	// Create record structure.
	rec := AdminData{
		Type: TypeAdminAuth,
		Data: key,
	}
	// Insert into the table.
	res := db.Create(&rec)
	if res.Error != nil {
		logrus.Error(ErrInsertRecordFail.Error() + " : " + res.Error.Error())
		return ErrInsertRecordFail
	}
	logrus.Info(fmt.Sprintf("Add access key and %d raws affected.", res.RowsAffected))
	return nil
}

// Query whether the access key exits.
func QueryHasAccessKey(key string) bool {
	// Empty key is invalid.
	if key == "" {
		logrus.Error(ErrInvalidAccessKey.Error())
		return false
	}
	// Set the target access key.
	rule := AdminData{
		Type: TypeAdminAuth,
		Data: key,
	}
	// Query in db.
	res := db.Where(&rule).First(&AdminData{})
	// if res.Error != nil {
	// 	logrus.Error(res.Error.Error())
	// }
	return res.Error == nil
}

// Insert an category.
func AddCategory(name string) (uint, error) {
	// Create record structure.
	rec := AdminData{
		Type: TypeAdminCategory,
		Data: name,
	}
	// Insert into the table.
	res := db.Create(&rec)
	if res.Error != nil {
		logrus.Error(ErrInsertRecordFail.Error() + " : " + res.Error.Error())
		return 0, ErrInsertRecordFail
	}
	logrus.Info(fmt.Sprintf("Add access key and %d raws affected.", res.RowsAffected))
	return rec.ID, nil
}

// Query whether the category exits by name.
func QueryCategoryIdByName(name string) (uint, error) {
	// Empty name is invalid.
	if name == "" {
		logrus.Error(ErrInvalidAccessKey.Error())
		return 0, ErrInvalidAccessKey
	}
	rec := AdminData{}
	// Set the target access key.
	rule := AdminData{
		Type: TypeAdminCategory,
		Data: name,
	}
	// Query in db.
	res := db.Where(&rule).First(&rec)
	if res.Error != nil {
		logrus.Error(res.Error.Error())
		return 0, res.Error
	}
	return rec.ID, nil
}
