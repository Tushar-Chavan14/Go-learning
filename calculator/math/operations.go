package operation

var Result int

func Addition(x int, y int) {
	Result = x + y
}

func Substraction(x int, y int) (int, int, int) {
	return x - y, x, y
}

func Multiplication(x int, y int) int {
	return x * y
}

func Divide(x int, y int) int {
	return x / y
}
