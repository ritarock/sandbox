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

func (t Task) FindId(id uint) Task {
	db := GormConnect()
	defer db.Close()

	var task Task
	db.Where("id = ?", id).Find(&task)
	return task
}

func (t Task) Create() {
	db := GormConnect()
	defer db.Close()

	var task = Task{Name: t.Name, Status: t.Status}
	db.Create(&task)
}

func (t Task) Get() *gorm.DB {
	db := GormConnect()
	defer db.Close()

	var task = Task{Name: t.Name}
	return db.First(&task)
}

func (t Task) Update() {
	db := GormConnect()
	defer db.Close()

	var task = Task{Name: t.Name}
	db.Updates(&task)
}

func (t Task) Delete() {
	db := GormConnect()
	defer db.Close()

	var task = Task{Model: gorm.Model{ID: t.ID}}
	db.Delete(&task)
}
