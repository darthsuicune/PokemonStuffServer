package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"encoding/json"
	"io/ioutil"
)

const driverName = "sqlite3"
const dbName = "./pokemon.db"

const createTableQuery = "CREATE TABLE IF NOT EXISTS pokemon (id integer, form integer, stamina integer, attack integer, defense integer," +
	"PRIMARY KEY(id, form))"

const checkTableHasContentQuery = "SELECT count(*) FROM pokemon"

func CreateDbIfNeeded() {
	db := OpenDb()
	defer db.Close()
	_, err := db.Exec(createTableQuery)
	checkError(err)
	if tableIsEmpty(db) {
		insertMons(db)
	}
}

func OpenDb() *sql.DB {
	db, err := sql.Open(driverName, dbName)
	checkError(err)
	return db
}

func tableIsEmpty(db *sql.DB) bool {
	rows, err := db.Query(checkTableHasContentQuery)
	checkError(err)
	defer rows.Close()

	rows.Next()
	var count int
	if err := rows.Scan(&count); err != nil {
		checkError(err)
	}
	return count == 0
}

func insertMons(db *sql.DB) {
	var dexItems Pokedex
	stats, err := ioutil.ReadFile("assets/mon_stats")
		checkError(err)
	err = json.Unmarshal(stats, &dexItems)
	checkError(err)

	//Start transaction
	tx, err := db.Begin()
	checkError(err)
	defer tx.Commit()

	for i := range dexItems.Stats  {
		//Insert all mons
		stmt, err := db.Prepare("INSERT INTO pokemon(id, form, attack, defense, stamina) values (?,?,?,?,?)")
		checkError(err)

		mon := dexItems.Stats[i]

		_, err = stmt.Exec(mon.Id, mon.Form, mon.Attack, mon.Defense, mon.Stamina)
		checkError(err)

	}
}

func checkError(e error) {
	if e != nil {
		println(e.Error())
		panic(e)
	}
}
