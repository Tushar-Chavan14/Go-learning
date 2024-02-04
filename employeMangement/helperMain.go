package main

import "fmt"

var (
	name, email string
)
var (
	empNo, age uint
	salary     uint64
)

var fieldToUpdate string

func getEmployeeData() {
	fmt.Println("Enter employee number : ")
	fmt.Scan(&empNo)
	fmt.Println("Enter employee name : ")
	fmt.Scan(&name)
	fmt.Println("Enter employee email : ")
	fmt.Scan(&email)
	fmt.Println("Enter employee age : ")
	fmt.Scan(&age)
	fmt.Println("Enter employee salary : ")
	fmt.Scan(&salary)
}

func getUpdateData() {
	fmt.Println("Enter employee number")
	fmt.Scan(&empNo)
	fmt.Println("Enter field to Update")
	fmt.Scan(&fieldToUpdate)
}
