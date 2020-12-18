package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(10)"`
}

func main() {
	router := gin.Default()

	dbInit()

	router.GET("/index", func(c *gin.Context) {
		users := dbGetAll()
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	router.Run()
}

func dbInit() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&User{})
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	PROTOCOL := "tcp(db:3306)"
	USER := "user"
	PASS := "password"
	DBNAME := "sample_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbGetAll() []User {
	db := gormConnect()

	defer db.Close()
	var users []User
	db.Find(&users)
	return users
}
