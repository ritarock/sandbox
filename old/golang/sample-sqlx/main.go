package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mithrandie/csvq-driver"
)

func main() {
	run()
}

var schemaPerson = "CREATE TABLE `person.csv` (first_name, last_name, email)"
var schemaPlace = "CREATE TABLE `place.csv` (country, city, telcode)"

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	Telcode string
}

func run() {
	db, err := sqlx.Connect("csvq", "./csvq")
	if err != nil {
		fmt.Println(err)
	}

	drop()
	db.MustExec(schemaPerson) // MustExec はエラー時に panic になる
	db.MustExec(schemaPlace)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO `person.csv` (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO `person.csv` (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO `place.csv` (country, city, telcode) VALUES (?, ?, ?)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO `place.csv` (country, telcode) VALUES (?, ?)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO `place.csv` (country, telcode) VALUES (?, ?)", "Singapore", "65")
	tx.NamedExec("INSERT INTO `person.csv` (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()

	people := []Person{}
	db.Select(&people, "SELECT * FROM `person.csv` ORDER BY first_name ASC")
	jason, john := people[0], people[1]
	fmt.Printf("%#v\n%#v\n", jason, john)

	jason = Person{}
	_ = db.Get(&jason, "SELECT * FROM `person.csv` Where first_name = ?", "Jason")
	fmt.Printf("%#v\n", jason)

	places := []Place{}
	err = db.Select(&places, "SELECT * FROM `place.csv` ORDER BY telcode ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	usa, singsing, honkers := places[0], places[1], places[2]
	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)

	place := Place{}
	rows, _ := db.Queryx("SELECT * FROM `place.csv`")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v\n", place)
	}

	_, _ = db.NamedExec(`INSERT INTO person (first_name, last_name, email) VALUES (:first, :last, :email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	rows, _ = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`,
		map[string]interface{}{"fn": "Bin"})

	rows, _ = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)

	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}
	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
		VALUES (:first_name, :last_name, :email)`, personStructs)

	personMaps := []map[string]interface{}{
		{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
		{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
		{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
		VALUES (:first_name, :last_name, :email)`, personMaps)
}

func drop() {
	os.RemoveAll("./csvq/person.csv")
	os.RemoveAll("./csvq/place.csv")
}
