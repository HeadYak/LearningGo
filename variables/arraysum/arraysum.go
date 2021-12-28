package main

import (
	"fmt"
)

func main() {

	var nums []int

	var num int

	for {
		fmt.Print("Enter number: ")

		fmt.Scan(&num)

		if num == 0 {
			break
		} else {
			nums = append(nums, num)
		}

	}
	fmt.Print(sum(nums))
}

// func add(x int, y int) int {
// 	return x + y
// }

func sum(array []int) int {
	var total int = 0
	for _, value := range array {
		total = total + value
	}
	return total
}
