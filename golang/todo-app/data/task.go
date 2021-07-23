package data

import "fmt"

type Task struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Detail    string `db:"detail" json:"detail"`
	Status    int    `db:"status" json:"status"`
	UserId    int    `db:"user_id" json:"user_id"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func TasksAll(user_id int) []Task {
	db := ConnectDb()
	defer db.Close()
	tasks := []Task{}
	err := db.Select(&tasks, "SELECT * FROM tasks WHERE user_id = ?", user_id)
	if err != nil {
		fmt.Println(err)
	}
	return tasks
}

func (task *Task) Create() {
	db := ConnectDb()
	defer db.Close()
	now := nowTime()
	query := `INSERT INTO tasks (name, detail, status, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	id, _ := db.MustExec(query, task.Name, task.Detail, 0, task.UserId, now, now).LastInsertId()
	task.ID = int(id)
	task.CreatedAt, task.UpdatedAt = now, now
}

func (task *Task) Read() {
	db := ConnectDb()
	defer db.Close()
	err := db.Get(task, "SELECT * FROM tasks WHERE id = ? AND user_id = ?", task.ID, task.UserId)
	if err != nil {
		fmt.Println(err)
	}
}
