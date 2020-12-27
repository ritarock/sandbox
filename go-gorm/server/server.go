package server

import (
	"fmt"
	"go-gorm/model"

	"github.com/gin-gonic/gin"
)

const VERSION = "v0"

func Run() {
	engine := gin.Default()
	model.DbInit()

	router := engine.Group("service")
	{
		api := router.Group(VERSION)
		{
			api.GET("/tasks", func(c *gin.Context) {
				task := model.Task{}
				c.JSON(200, gin.H{
					"tasks": task.GetAll(),
				})
			})

			api.POST("/tasks", func(c *gin.Context) {
				task := model.Task{}
				err := c.ShouldBindJSON(&task)
				if err != nil {
					fmt.Println(err)
				}
				task.Create()
				c.Redirect(302,
					"/service/"+VERSION+"/tasks")
			})

			api.GET("tasks/:id", func(c *gin.Context) {
				var t model.Task
				c.JSON(200, gin.H{
					"task": t.Get(c.Param("id")),
				})
			})

			api.PUT("tasks/:id", func(c *gin.Context) {
				task := model.Task{}
				err := c.ShouldBindJSON(&task)
				if err != nil {
					fmt.Println(err)
				}
				task.Update(c.Param("id"))
				c.Redirect(302,
					"/service/"+VERSION+"/tasks")
			})

			api.DELETE("tasks/:id", func(c *gin.Context) {
				task := model.Task{}
				task.Delete(c.Param("id"))
				c.Redirect(302,
					"/service/"+VERSION+"/tasks")
			})
		}
	}

	engine.Run()
}
