package model

import (
	"database/sql"
	"os"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "todo"
	dbPass := os.Getenv("MYSQL_TODO_PASSWORD")
	dbName := "tododb"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
