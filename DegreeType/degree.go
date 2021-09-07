package main

import "fmt"

//Create a new type degree
//This type is  a slice of string
type degree []string

func (d degree) print() {
	for i, degrees := range d {
		fmt.Println(i, degrees)
	}
}
