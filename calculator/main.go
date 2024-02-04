package main

import (
	"fmt"

	operation "calculator/math"
)

func main() {
	greet()
	options()

	var choice int

	for {
		fmt.Print("Enter your choice : ")
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		switch choice {
		case 1:
			getInputs()
			operation.Addition(firstNumber, secondNumber)
			println("result :", operation.Result)
			options()
		case 2:
			getInputs()
			difference, x, y := operation.Substraction(firstNumber, secondNumber)
			fmt.Printf("The result %d for %d and %d\n", difference, x, y)
			options()
		case 3:
			getInputs()
			product := operation.Multiplication(firstNumber, secondNumber)
			println("result :", product)
			options()
		case 4:
			getInputs()
			quotient := operation.Divide(firstNumber, secondNumber)
			println("result :", quotient)
			options()

		default:
			fmt.Println("Enter any specified option")
			options()
		}
	}

	fmt.Println("Thank You! Have fun")
}
