package main

import "fmt"

var (
	firstNumber  int
	secondNumber int
)

func greet() {
	fmt.Println("Welcome")
	fmt.Println("Cmd-Calculator")
}

func options() {
	fmt.Println("Please opt an option")
	fmt.Println("1: addition")
	fmt.Println("2: substraction")
	fmt.Println("3: multiplication")
	fmt.Println("4: divide")
	fmt.Println("0: Exit")
}

func getInputs() {
	fmt.Println("Enter two numbers space seprated :")
	fmt.Scan(&firstNumber, &secondNumber)
}
