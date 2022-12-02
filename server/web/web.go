package web

import (
	"IroChest/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var engine *gin.Engine

// Initialization.
func Init() {
	engine = gin.Default()
	r := engine // Alias

	api := r.Group("/api")
	{
		api.POST("/note/new", CreateNewNote)
		api.POST("/note/list", ListNoteByCategory) // The special situation of `ListNoteByCategory`
		api.POST("/note/list/category", ListNoteByCategory)
	}
}

func Run() {
	r := engine // Alias
	logrus.Info(config.Conf.Port)
	if err := r.Run(":" + config.Conf.Port); err != nil {
		logrus.Fatal(ErrRunEngineFail)
	}
}
