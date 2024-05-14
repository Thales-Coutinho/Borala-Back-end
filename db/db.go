package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DataBaseConection() *sql.DB {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// type Config struct {
	// 	DatabaseHost     string    `env:"DATABASE_HOST,required"`
	// 	DatabasePort     string    `env:"DATABASE_PORT,required"`
	// 	DatabaseUsername string    `env:"DATABASE_USERNAME,required"`
	// 	DatabasePassword string    `env:"DATABASE_PASSWORD,required"`
	// 	DatabaseDBName   string    `env:"DATABASE_DBNAME,required"`
	// }

	conection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
