package data

import (
	"fmt"
)

type User struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func UsersAll() []User {
	db := ConnectDb()
	defer db.Close()
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func (user *User) Create() {
	db := ConnectDb()
	defer db.Close()
	now := nowTime()
	query := `INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, ?, ?)`
	id, _ := db.MustExec(query, user.Name, user.Email, now, now).LastInsertId()
	user.ID = int(id)
	user.CreatedAt, user.UpdatedAt = now, now
}

func (user *User) Read() {
	db := ConnectDb()
	defer db.Close()
	err := db.Get(user, "SELECT * FROM users WHERE id = ?", user.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func (user *User) Update(newUser User) {
	db := ConnectDb()
	defer db.Close()
	now := nowTime()
	if newUser.Name == "" {
		newUser.Name = user.Name
	}
	if newUser.Email == "" {
		newUser.Email = user.Email
	}
	newUser.UpdatedAt = now
	query := `UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?`
	db.MustExec(query, newUser.Name, newUser.Email, now, user.ID)
}

func (user *User) Delete() {
	db := ConnectDb()
	defer db.Close()
	query := `DELETE FROM users WHERE id = ?`
	db.MustExec(query, user.ID)
}
