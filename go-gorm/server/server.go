package server

import (
	"go-gorm/model"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	model.DbInit()

	router.GET("/index", func(c *gin.Context) {
		tasks := model.GetTaskAll()
		c.JSON(200, gin.H{
			"tasks": tasks,
		})
	})

	router.GET("/create", func(c *gin.Context) {
		model.CreateTask()
	})

	router.Run()
}
