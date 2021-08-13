package main

import "fmt"

func main() {
	//We declare a slice of numbers
	numbers := []int{1, 2, 3}
	//We add a number using the function append(slice, newItem)
	numbers = append(numbers, 4)

	//for loop to iterate on the slice previously created
	for i, numbers := range numbers {
		fmt.Println(i, numbers)
	}

	//We declare an array of strings
	signs := [3]string{"Gemini", "Leo", "Cancer"}
	//Print it
	for i, signs := range signs {
		fmt.Println(i, signs)
	}
}
