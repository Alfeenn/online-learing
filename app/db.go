package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Alfeenn/online-learning/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Unable to load env")
	}
	statement := helper.SQLStatement(
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)
	log.Print(statement)
	db, err := sql.Open("mysql", statement)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	return db
}
