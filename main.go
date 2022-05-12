package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid Usage!")
		fmt.Println("COMMANDS:")
		fmt.Println("\tadd\t add a new record")
		fmt.Println("\tquery\t query a record")
		fmt.Println("\tupdate\t update a record")
		fmt.Println("\tremove\t remove a record")
		fmt.Println("\tlist\t list all the records")
		os.Exit(1)
	}
	var person Person
	switch os.Args[1] {
	case "add":
		person.Name = os.Args[2]
		person.Age, _ = strconv.Atoi(os.Args[3])
		person.Insert()
	case "query":
		id, _ := strconv.Atoi(os.Args[2])
		person = GetElementByID(id)
		fmt.Println("Name: ", person.Name)
		fmt.Println("Age: ", person.Age)
	case "update":
		var new_person Person
		update_id, _ := strconv.Atoi(os.Args[2])
		new_person.Name = os.Args[3]
		new_person.Age, _ = strconv.Atoi(os.Args[4])
		Update(update_id, &new_person)
	case "remove":
		Delete(os.Args[2])
	case "list":
		ListAllElements()
	default:
		fmt.Println("Invalid command!")
		os.Exit(1)
	}
}
