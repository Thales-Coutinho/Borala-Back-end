package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DataBaseConection() *sql.DB {
	conection := "root:123456@tcp(localhost:3306)/Borala"
	db, err := sql.Open("mysql", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
