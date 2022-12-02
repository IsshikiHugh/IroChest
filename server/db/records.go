package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * `db/records.go`												 *
 * Be used to do the CURD option for note records.				 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// Add a new note record.
func AddNewNote(note *Note) (uint, error) {
	cid, err := QueryCategoryIdByName(note.Category)
	if err != nil {
		cid, err = AddCategory(note.Category)
		if err != nil {
			return 0, err
		}
	}

	record := Records{
		Category: cid,
		Data:     note.Data,
	}

	return insertRecord(&record)
}

// Insert a record into record table.
func insertRecord(rec *Records) (uint, error) {
	res := db.Create(&rec)
	if res.Error != nil {
		logrus.Error(ErrInsertRecordFail.Error() + " : " + res.Error.Error())
		return 0, ErrInsertRecordFail
	}
	logrus.Info(fmt.Sprintf("Add access key and %d raws affected.", res.RowsAffected))
	return rec.ID, nil
}
