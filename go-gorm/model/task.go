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

func (t Task) Get(id string) Task {
	db := GormConnect()
	defer db.Close()

	var task Task
	db.Where("id = ?", id).Find(&task)
	return task
}

func (t Task) Update(id string) {
	db := GormConnect()
	defer db.Close()

	var task Task
	db.Where("id = ?", id).Find(&task)
	task.Name = t.Name
	task.Status = t.Status
	db.Save(&task)
}

func (t Task) Delete(id string) {
	db := GormConnect()
	defer db.Close()

	var task Task
	db.Where("id = ?", id).Find(&task)
	db.Delete(&task)
}
