package data

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

var dropSchema = `
DROP TABLE IF EXISTS users;
`
var schema = `
CREATE TABLE users (
	id int,
	name string,
	email string,
	created_at int,
	updated_at int
);
`

const (
	DBMS     = "mysql"
	PROTOCOL = "tcp(db:3306)"
	USER     = "user"
	PASS     = "password"
	DBNAME   = "app"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
)

func buildSqlx() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		fmt.Println("Not ready")
		time.Sleep(time.Second)
		return buildSqlx()
	}
	fmt.Println("Success connect")
	return db
}

func connectSqlx() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success connect")
	return db
}

func init() {
	db := buildSqlx()
	db.MustExec(dropSchema)
	db.MustExec(schema)
}
