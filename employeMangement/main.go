package main

import (
	"employee/employee"
	"fmt"
)

type Employee struct {
	empNo  uint
	name   string
	email  string
	age    uint
	salary uint64
}

func greet() {
	fmt.Println("welcome admin")
	fmt.Println("employee management")
}

func options() {
	fmt.Println("Please opt an option")
	fmt.Println("1: add emplyee")
	fmt.Println("2: view employee")
	fmt.Println("3: list employee")
	fmt.Println("4: update employee")
	fmt.Println("5: remove employee")
	fmt.Println("0: Exit")
}

func main() {
	var choice uint

	greet()

	for {
		options()

		fmt.Println("Enter your choice : ")
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		switch choice {
		case 1:
			getEmployeeData()
			employee.AddEmployee(empNo, name, email, age, salary)
			fmt.Println("employee added")
		case 2:
			fmt.Println("Enter employee number")
			fmt.Scan(&empNo)
			employee.ViewEmployee(empNo)
		case 3:
			employees := employee.ListEmployee()
			fmt.Println(employees)
		case 4:
			getUpdateData()
			employee.UpdateEmployee(empNo, fieldToUpdate)
			fmt.Println("employee updated")
		case 5:
			employee.RemoveEmployee()
			employees := employee.ListEmployee()
			fmt.Println(employees)

		default:
			fmt.Println("Enter any of the specified option")
		}

	}

}
