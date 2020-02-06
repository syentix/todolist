package main

import (
	"database/sql"
	"fmt"
	"os"

	"./config"
	"./database"
	"./input"
	"./telegram"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) > 1 {
		// Reading Config-File.
		err := config.ReadConfig()
		if err != nil {
			panic(err.Error())
		}
		// CLI-Mode
		if os.Args[1] == "-c" {
			database.DBCon, err = sql.Open("mysql", config.Creds)
			if err != nil {
				panic(err.Error())
			}
			// Opening Inputreader.
			input.StartReader()
			defer database.DBCon.Close()
		}
		if os.Args[1] == "-tele" {
			telegram.StartBot()
		}
		os.Exit(0)
	} else {
		fmt.Println("No flag given, what da hell am I supposed to do now?")
		defer os.Exit(0)
	}

}
