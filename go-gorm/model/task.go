package model

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	gorm.Model
	Name   string
	Status int
}

func GetTaskAll() []Task {
	db := GormConnect()
	defer db.Close()

	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func CreateTask() {
	db := GormConnect()
	defer db.Close()

	var task = Task{Name: "test", Status: 0}
	fmt.Println(task)
	db.Create(&task)
}
