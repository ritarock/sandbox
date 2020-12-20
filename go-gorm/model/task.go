package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	gorm.Model
	Name   string
	Status int
}

func (t Task) GetAll() []Task {
	db := GormConnect()
	defer db.Close()

	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func (t Task) Create() {
	db := GormConnect()
	defer db.Close()

	var task = Task{Name: t.Name, Status: t.Status}
	db.Create(&task)
}
