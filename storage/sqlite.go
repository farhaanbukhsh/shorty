package storage

import (
	"database/sql"
	"fmt"

	"github.com/farhaanbukhsh/shorty/encoder"
	_ "github.com/mattn/go-sqlite3"
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
		code VARCHAR(64) NOT NULL UNIQUE);
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
	var slugExists int

	err := conn.Db.QueryRow("SELECT COUNT(*) FROM shorty where code=$1", slug).Scan(&slugExists)
	if err != nil {
		return "", fmt.Errorf("Slug exists check failed, %s", err)
	}

	statement, err := conn.Db.Prepare("INSERT INTO shorty (url, code) VALUES (?, ?)")
	if err != nil {
		return "", fmt.Errorf("Insert statement is wrong, %s", err)
	}

	if slugExists != 0 {
		return "", fmt.Errorf("Slug Already Exists")
	}

	if slug != "" {
		code = slug
	} else {
		code, err = generateURLCode(conn)
		if err != nil {
			return "", fmt.Errorf("Cannot generate code")
		}
	}

	_, err = statement.Exec(url, code)
	if err != nil {
		return "", fmt.Errorf("Cannot add code to the database")
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

func generateURLCode(conn sqliteConnetion) (string, error) {
	var rows int
	err := conn.Db.QueryRow("SELECT COUNT(*) FROM shorty").Scan(&rows)
	if err != nil {
		return "", fmt.Errorf("Row query didn't execute")
	}
	code := encoder.URLCodeGenerator(rows)
	return code, nil
}
