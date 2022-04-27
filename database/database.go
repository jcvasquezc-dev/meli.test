package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var sqliteDbFileName string = "xmen.db" // default database

func Initialize() {
	db := GetConnection()

	if _, err := db.Exec(QRY_CREATE_TABLE_DNA); err != nil {
		panic(err)
	}

	db.Close()
}

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", sqliteDbFileName)

	if err != nil {
		panic(err)
	}

	return db
}

func InitializeForTest() {
	const sqliteDbFileNameForTest string = "test-xmen.db"
	_ = os.Remove(sqliteDbFileNameForTest)

	sqliteDbFileName = sqliteDbFileNameForTest
	db := GetConnection()

	if _, err := db.Exec(QRY_CREATE_TABLE_DNA); err != nil {
		panic(err)
	}

	db.Close()
}
