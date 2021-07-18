package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var count int
var Db *sql.DB

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Tasks     []Task
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"create_at"`
}

type Task struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"create_at"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	PROTOCOL := "tcp(db:3306)"
	USER := "user"
	PASS := "password"
	DBNAME := "app"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	var err error
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println("Not ready")
		time.Sleep(time.Second)
		count++
		if count > 30 {
			panic(err.Error())
		}
		return gormConnect()
	}
	fmt.Println("Success connect")
	return db
}

func init() {
	fmt.Println("Started")
	db := gormConnect()
	db.AutoMigrate(&User{}, &Task{})
}
