package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../database"
)

type ToDo struct {
	ID      int    `json:"id"`
	Text    string `json:"todo"`
	Date    string `json:"date"`
	Checked bool   `json:"completed"`
}

func StartReader() {
	reader := bufio.NewReader(os.Stdin)
	printToDoList()
	for {
		fmt.Print("->")
		// Read Input.
		in, _ := reader.ReadString('\n')
		in = strings.Replace(in, "\n", "", -1)
		inArray := strings.Split(in, " ")

		// Retrieve Command
		command := inArray[0]

		// Delete Command from string splice, in order to put message together.
		inArray = append(inArray[:0], inArray[1:]...)
		todo := strings.Join(inArray, " ")

		// Decide on Case...
		switch command {
		case "add":
			fmt.Println("==>Adding: " + todo + " <==")
			addToDB(todo)
			printToDoList()
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(100)
		case "check":
			checkID, err := strconv.Atoi(todo)
			if err != nil {
				fmt.Println("Try Again! Usage is: check <ID>")
			}
			check(checkID)
			printToDoList()
		case "print":
			printToDoList()
		default:
			fmt.Println("No Such Command.")
			fmt.Println("List of valid commands: \n +add\n +check\n +exit\n +print")
		}
	}
}

func addToDB(msg string) {
	query := fmt.Sprintf("INSERT INTO todos (todo) VALUES ('%s')", msg)
	insert, err := database.DBCon.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func printToDoList() {
	fmt.Println()
	var myToDos []ToDo
	results, err := database.DBCon.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	for results.Next() {
		var todo ToDo

		err = results.Scan(&todo.ID, &todo.Text, &todo.Date, &todo.Checked)
		if err != nil {
			panic(err.Error())
		}

		myToDos = append(myToDos, todo)
	}

	fmt.Println("Your ToDos! ================")
	for _, todo := range myToDos {
		var checkmark string
		if todo.Checked == true {
			checkmark = "\u2713"
		} else {
			checkmark = "\u2716"
		}
		fmt.Printf("%d) %s | Created: %s | %s\n", todo.ID, todo.Text, todo.Date, checkmark)
	}
	fmt.Println()
}

func check(checkID int) {
	// Check if ID exits!
	query := fmt.Sprintf("SELECT id FROM todos WHERE id=%d", checkID)
	selectQ, _ := database.DBCon.Query(query)
	if selectQ.Next() == false {
		fmt.Println("ERROR: ID doesn't exist. Try again!")
		return
	}
	defer selectQ.Close()
	// Updating.
	query = fmt.Sprintf("UPDATE todos SET completed=TRUE WHERE id=%d", checkID)
	update, err := database.DBCon.Query(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("== Good Job! ==")
	defer update.Close()
}
