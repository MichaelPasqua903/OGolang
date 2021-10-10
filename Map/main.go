package main

import "fmt"

func main() {
	//declaring a map in golang
	//            key   values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// Add a new key-value pair in a map
	colors["white"] = "#ffffff"

	//fmt.Println(colors)
	printMap(colors)
	//Another way of declaring a map in GO
	//var nations map[string]string

	//fmt.Println(nations)

	//third way of delcaring a map in GO
	//using the built-in function make
	//languages := make(map[string]string)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for " + color + " is " + hex)
	}
}
