package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Loading Sqlite driver
)

// SqliteConnetion responsible to connect to the database
type sqliteConnetion struct {
	Db *sql.DB
}

//New connection with the database
func New(databaseName string) (Service, error) {
	database, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		panic(err)
	}
	createTable := `CREATE TABLE IF NOT EXISTS 
		shorty (uid INTEGER PRIMARY KEY AUTOINCREMENT, url VARCHAR(256) NOT NULL,
		code VARCHAR(64) NOT NULL);
	`

	_, err = database.Exec(createTable)

	if err != nil {
		return nil, fmt.Errorf("Cannot create table, %s", err)
	}
	return &sqliteConnetion{database}, nil
}

// Save method helps to save the url and get back a short code
func (conn sqliteConnetion) Save(url string, slug string) (string, error) {
	var code string
	statement, err := conn.Db.Prepare("INSERT INTO shorty (url, code) VALUES (?, ?)")
	if err != nil {
		return "", fmt.Errorf("Insert statement is wrong, %s", err)
	}

	if slug != "" {
		statement.Exec(url, slug)
		code = slug
	} else {
		code = "bleh"
		statement.Exec(url, code)
	}
	return code, nil
}

func (conn sqliteConnetion) Load(code string) (string, error) {
	var url string
	err := conn.Db.QueryRow("SELECT url FROM shorty where code=$1 limit 1", code).Scan(&url)
	if err != nil {
		return "", fmt.Errorf("Url retrieval problem, %s", err)
	}
	return url, nil
}

func (conn sqliteConnetion) Close() error {
	err := conn.Db.Close()
	if err != nil {
		return fmt.Errorf("Cannot close connection, %s", err)
	}
	return nil
}
