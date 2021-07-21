package data

import (
	"fmt"
	"time"
)

type User struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func GetUsers() (users []User, err error) {
	db := connectSqlx()
	users = []User{}
	err = db.Get(&users, "SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (user *User) Create() (err error) {
	db := connectSqlx()
	tx := db.MustBegin()
	tx.MustExec(
		"INSERT INTO (name, email, created_at, updated_up) VALUES ($1, $2, $3, $4)",
		user.Name, user.Email, time.Now(), time.Now())
	tx.Commit()
	return
}
