package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

	return InsertRecord(&record)
}

// Insert a record into record table.
func InsertRecord(rec *Records) (uint, error) {
	res := db.Create(&rec)
	if res.Error != nil {
		logrus.Error(ErrInsertRecordFail.Error() + " : " + res.Error.Error())
		return 0, ErrInsertRecordFail
	}
	logrus.Info(fmt.Sprintf("Add access key and %d raws affected.", res.RowsAffected))
	return rec.ID, nil
}

// Query notes by category within limit.
// If cname is "", then means every category is available.
// If lim is 0, then means every records will be queried.
func QueryNotesByCategoryAndLimit(cname string, lim uint) ([]Note, error) {
	var (
		cid uint
		err error
		res *gorm.DB
	)

	// Set query rule
	if cname == "" {
		cid = 0
	} else {
		cid, err = QueryCategoryIdByName(cname)
		if err != nil {
			return nil, err
		}
	}
	rule := Records{
		Category: cid, // 0 will be seen as no rule.
	}
	// Query.
	recs := []Records{}
	if lim == 0 {
		res = db.Where(&rule).Find(&recs)
	} else {
		res = db.Where(&rule).Limit(int(lim)).Find(&recs)
	}
	if res.Error != nil {
		logrus.Error(ErrQueryRecordFail.Error() + " : " + res.Error.Error())
		return nil, ErrQueryRecordFail
	}
	ret := []Note{}
	for _, v := range recs {
		ret = append(ret, Note{
			Category: cname,
			Data:     v.Data,
			Created:  v.Created,
		})
	}
	return ret, nil
}
