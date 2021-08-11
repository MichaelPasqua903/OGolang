package main

import "fmt"

func main() {
	//We can declare (and initialize) a varible as follows:
	//var card string = "The ace of spades"
	//or
	card := "The ace of spade"
	//If we want to modify the value of the variable created above we use the = operator
	card = "The ace of something"
	//print card
	fmt.Println(card)
}
