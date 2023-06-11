package main

import "fmt"

// This code creates a list of integers from 0 to 10 using a for loop with the append function.
//Then, it iterates over the list using a range loop and prints whether each number is even or odd.
func main() {
	var list []int
	for i := 0; i <= 10; i++ {
		list = append(list, i)
	}

	for _, num := range list {
		isEven := "even"
		if num%2 != 0 {
			isEven = "odd"
		}
		fmt.Printf("%d is %s\n", num, isEven)
	}
}