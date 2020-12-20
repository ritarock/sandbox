package server

import (
	"go-gorm/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	model.DbInit()

	router := engine.Group("service")
	{
		v0 := router.Group("v0")
		{
			v0.GET("/index", func(c *gin.Context) {
				task := model.Task{}
				c.JSON(200, gin.H{
					"tasks": task.GetAll(),
				})
			})

			v0.POST("/new", func(c *gin.Context) {
				var t model.Task
				t.Name = c.PostForm("Name")
				t.Status, _ = strconv.Atoi(c.PostForm("Status"))

				task := model.Task{Name: t.Name, Status: t.Status}
				task.Create()
				c.Redirect(302, "/service/v0/index")
			})
		}
	}

	engine.Run()
}
