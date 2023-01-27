package go_kill_you

import (
	"fmt"
)

func KillYou(name string) string {
	return "I really want to kill you " + name
}

func PersonId() {
	var name string
	var age string
	var phone int
	fmt.Println("Name: ")
	fmt.Scan(&name)
	fmt.Println("Age: ")
	fmt.Scan(&age)
	fmt.Println("HP: ")
	fmt.Scan(&phone)
	fmt.Println("--------------------")

	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old")
	fmt.Println("My phone number is", phone)

}
