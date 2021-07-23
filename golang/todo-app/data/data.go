package data

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DBMS     = "mysql"
	PROTOCOL = "tcp(db:3306)"
	USER     = "user"
	PASS     = "password"
	DBNAME   = "app"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
)

func initDB() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		fmt.Println("Not ready")
		time.Sleep(time.Second)
		return initDB()
	}
	fmt.Println("Sucess connect")
	return db
}

func ConnectDb() *sqlx.DB {
	db, err := sqlx.Connect(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func init() {
	fmt.Println("Started")
	initDB()
}

const FORMAT = "2006-01-02 15:04:05"

func nowTime() string {
	return time.Now().Format(FORMAT)
}
