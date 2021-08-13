package main

import "fmt"

func main() {
	//call a function defined in this file to initialize the following variable
	newQuote := quote()
	fmt.Println(newQuote)
	//I can call a function from another file without  any other import because the two files share the same package
	printSomething()

}

func quote() string {
	return "This is quote"
}
