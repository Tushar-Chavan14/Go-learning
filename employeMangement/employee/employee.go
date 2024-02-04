package employee

import (
	"fmt"
)

type Employee struct {
	empNo  uint
	name   string
	email  string
	age    uint
	salary uint64
}

var employees = make([]Employee, 0)

func AddEmployee(empNo uint, name string, email string, age uint, sal uint64) {
	var employeeData = Employee{
		empNo:  empNo,
		name:   name,
		email:  email,
		age:    age,
		salary: sal,
	}

	employees = append(employees, employeeData)
	return
}

func ViewEmployee(empNo uint) {
	emplyee, isExist := Find(employees, func(e Employee) bool {
		return empNo == e.empNo
	})

	if !isExist {
		fmt.Println("No employee found with that employee number")
		return
	}

	fmt.Println(emplyee)
	return
}

func ListEmployee() []Employee {
	return employees
}

func RemoveEmployee() {

	var empNo uint

	fmt.Println("Enter employee number")
	fmt.Scan(&empNo)

	employee := Filter(employees, func(e Employee) bool {
		return empNo != e.empNo
	})

	employees = employee

}

func UpdateEmployee(empNo uint, fieldToUpdate string) {

	var (
		name, email string
		age         uint
		sal         uint64
	)

	_, isExist := Find(employees, func(e Employee) bool {
		return empNo == e.empNo
	})

	if !isExist {
		fmt.Println("No employee found with that employee number")
		return
	}

	for i, elem := range employees {
		if empNo == elem.empNo {
			switch fieldToUpdate {
			case "name":
				fmt.Println("Enter name: ")
				fmt.Scan(&name)
				employees[i].name = name
			case "age":
				fmt.Println("Enter age: ")
				fmt.Scan(&age)
				employees[i].age = age
			case "email":
				fmt.Println("Enter email: ")
				fmt.Scan(&email)
				employees[i].email = email
			case "sal":
				fmt.Println("Enter salary: ")
				fmt.Scan(&sal)
				employees[i].salary = sal
			default:
				fmt.Println("unkown field")
				return
			}
		}
	}

	return

}
