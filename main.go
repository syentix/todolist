package main

import (
	"database/sql"
	"fmt"
	"os"

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
	if len(os.Args) > 1 {
		if os.Args[1] == "-c" {
			input.StartReader()
		} else if os.Args[1] == "-tele" {
			// TODO: Call Telegram-Bot.
		}
	} else {
		fmt.Println("No flag given, what da hell am I supposed to do now?")
		defer os.Exit(0)
	}

	defer database.DBCon.Close()
}
