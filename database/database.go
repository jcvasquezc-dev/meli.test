package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "xmen.db"

func Initialize() {
	db := GetConnection()

	if _, err := db.Exec(QRY_CREATE_TABLE_DNA); err != nil {
		panic(err)
	}

	db.Close()
}

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		panic(err)
	}

	return db
}
