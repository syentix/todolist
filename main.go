package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"./config"
	"./database"
	"./input"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	database.DBCon, err = sql.Open("mysql", config.Creds)
	if err != nil {
		panic(err.Error())
	}

	// Opening Inputreader.
	input.StartReader()

	defer database.DBCon.Close()
}
