package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Mariadb() *sql.DB {

	godotenv.Load(".env")
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", os.Getenv("MARIADB"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to MariaDB!")
	return db
}