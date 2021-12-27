package main

import (
	"fmt"
)

func main() {

	var num1 int
	var num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Println(add(num1, num2))

}

func add(x int, y int) int {
	return x + y
}
