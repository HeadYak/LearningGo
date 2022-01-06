package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var nums []int

	var input string

	for {
		fmt.Print("Enter number: ")

		_, err := fmt.Scanln(&input)

		if err != nil && err.Error() == "unexpected newline" {
			fmt.Println("")
			break
		} else {

			// fmt.Println(input)

			if num, err := strconv.Atoi(input); err == nil {
				nums = append(nums, num)
			} else {
				fmt.Println("Couldnt convert input to int")
			}
		}

	}

	if len(nums) != 0 {
		fmt.Println("Sum:", sum(nums))
		fmt.Println("Len:", length(nums))
		fmt.Println("Array:", nums)
		fmt.Println("Sorted array:", sortarray(nums))
	} else {
		fmt.Println("Empty array")
	}
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

func length(array []int) int {
	return len(array)
}

func sortarray(array []int) []int {
	var sorted []int = array
	// sorted = sort.Ints(array)
	sort.Ints(sorted)
	return sorted
}
