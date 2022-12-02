package web

import (
	"IroChest/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateNewNote(c *gin.Context) {
	req := CreateNewNoteReq{}
	// Check valid.
	if err := c.Bind(&req); err != nil {
		logrus.Error(ErrBindBodyFail.Error() + " : " + err.Error())
		c.JSON(http.StatusOK, ErrResp{
			ErrMsg: ErrBindBodyFail.Error(),
		})
		return
	}
	// Auth
	if !db.QueryHasAccessKey(req.AccessKey) {
		logrus.Error(ErrAuthFail.Error())
		c.JSON(http.StatusOK, ErrResp{
			ErrMsg: ErrAuthFail.Error(),
		})
		return
	}
	// Insert
	note := db.Note{
		Category: req.Category,
		Data:     req.Data,
	}
	db.AddNewNote(&note)
	c.JSON(http.StatusOK, SuccessResp{
		ErrMsg: "",
		Data:   note,
	})
	return
}
