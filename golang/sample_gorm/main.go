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

type Post struct {
	ID        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
	UpdatedAt time.Time
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

func main() {
	fmt.Println("Started")
	db := gormConnect()
	db.AutoMigrate(&Post{}, &Comment{})
	fmt.Println("Finished database migrate")

	post := Post{Content: "Hello", Author: "user1"}
	fmt.Println(post) // {0 Hello user1 [] 0001-01-01 00:00:00 +0000 UTC}

	db.Create(&post)
	fmt.Print(post) // {1 Hello user1 [] 2021-07-17 03:35:06.80860259 +0000 UTC m=+10.096848216}
	fmt.Println("Created post")

	comment := Comment{Content: "good", Author: "user2"}
	db.Model(&post).Association("Comments").Append(comment)
	fmt.Println("Created comment")

	var readPost Post
	db.Where("author = ?", "user1").First(&readPost)

	var comments []Comment
	db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0]) // {1 good user2 1 2021-07-17 03:35:07 +0000 UTC}
}
