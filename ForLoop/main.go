package main

import "fmt"

func main() {

	fmt.Println("Print the first 10 numbers")
	numbers := 10
	//Use the for loop to print the first 10 numbers
	for i := 0; i < numbers; i++ {
		println(i)
	}

	fmt.Println("Print the elements of an array")
	//Define and initialize an array
	points := []int{1, 2, 43, 45, 6}

	for i, points := range points {
		println(i, points)
	}

}
